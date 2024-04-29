package jobs

import (
	"context"
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/ent/feed"
	"github.com/ITK13201/rss-generator/ent/feeditem"
	"github.com/ITK13201/rss-generator/internal/db"
	"github.com/ITK13201/rss-generator/internal/rss"
	"github.com/ITK13201/rss-generator/internal/scraper"
	"github.com/sirupsen/logrus"
)

type UpdateFeedsJob interface {
	Run(ctx context.Context)
}

type updateFeedsJob struct {
	cfg       *domain.Config
	logger    *logrus.Logger
	sqlClient *ent.Client
}

func NewUpdateFeedsJob(cfg *domain.Config, logger *logrus.Logger, sqlClient *ent.Client) UpdateFeedsJob {
	return &updateFeedsJob{
		cfg:       cfg,
		logger:    logger,
		sqlClient: sqlClient,
	}
}

func (ufj *updateFeedsJob) Run(ctx context.Context) {
	ufj.logger.WithFields(logrus.Fields{
		"job": "update_feeds",
	}).Info("job started")

	sites, err := ufj.sqlClient.Site.Query().
		WithFeeds().
		WithScrapingSettings().
		All(ctx)
	if err != nil {
		ufj.logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("failed to site query")
	}

	for _, s := range sites {
		ufj.logger.WithFields(logrus.Fields{
			"site": s.Slug,
		}).Info("updating feed")

		if len(s.Edges.Feeds) < 1 {
			ufj.logger.WithFields(logrus.Fields{
				"site": s.Slug,
			}).Info("no feed found")
			continue
		}

		siteScraper := scraper.NewScraper(ufj.cfg, ufj.logger)
		newFeed, err := siteScraper.FetchFeedElements(
			s.URL,
			s.EnableJsRendering,
			domain.ConvertScrapingSettingFromModelToDomain(s.Edges.ScrapingSettings[0]))
		if err != nil {
			ufj.logger.WithFields(logrus.Fields{
				"site":  s.Slug,
				"error": err.Error(),
			})
		}

		oldFeed := domain.ConvertFeedFromModelToDomain(s.Edges.Feeds[0])

		rssUtil := rss.NewRssUtil(ufj.cfg, ufj.logger)
		updatedFeed := rssUtil.Update(oldFeed, newFeed)

		if err = db.WithTx(ctx, ufj.sqlClient, func(tx *ent.Tx) error {
			feedModel := s.Edges.Feeds[0]
			feedModel, err = tx.Feed.UpdateOne(feedModel).
				SetTitle(updatedFeed.Title).
				SetDescription(updatedFeed.Description).
				SetLink(updatedFeed.Link).
				SetPublishedAt(updatedFeed.PublishedAt).
				Save(ctx)
			if err != nil {
				return err
			}
			_, err = tx.FeedItem.Delete().
				Where(feeditem.HasFeedWith(feed.IDEQ(feedModel.ID))).
				Exec(ctx)
			if err != nil {
				return err
			}
			err = tx.FeedItem.MapCreateBulk(updatedFeed.Items, func(fic *ent.FeedItemCreate, i int) {
				fic.SetFeed(feedModel).
					SetTitle(updatedFeed.Items[i].Title).
					SetDescription(updatedFeed.Items[i].Description).
					SetLink(*updatedFeed.Items[i].Link).
					SetPublishedAt(updatedFeed.PublishedAt)
			}).Exec(ctx)
			if err != nil {
				return err
			}
			return nil
		}); err != nil {
			ufj.logger.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Fatal("failed to update feed")
		}
		ufj.logger.WithFields(logrus.Fields{
			"site": s.Slug,
		}).Info("updated feed")
	}

	ufj.logger.WithFields(logrus.Fields{
		"job": "update_feeds",
	}).Info("job finished")
}
