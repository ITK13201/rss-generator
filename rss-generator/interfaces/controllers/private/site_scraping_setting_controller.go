package private

import (
	"net/http"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/ent/feed"
	"github.com/ITK13201/rss-generator/ent/feeditem"
	"github.com/ITK13201/rss-generator/ent/scrapingsetting"
	"github.com/ITK13201/rss-generator/ent/site"
	"github.com/ITK13201/rss-generator/interfaces/interactors/private"
	"github.com/ITK13201/rss-generator/internal/db"
	"github.com/ITK13201/rss-generator/internal/rest"
	"github.com/ITK13201/rss-generator/internal/rss"
	"github.com/ITK13201/rss-generator/internal/scraper"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SiteScrapingSettingController interface {
	GetBySiteSlug(ctx *gin.Context)
	Upsert(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type siteScrapingSettingController struct {
	cfg                           *domain.Config
	logger                        *logrus.Logger
	sqlClient                     *ent.Client
	siteScrapingSettingInteractor private.SiteScrapingSettingInteractor
}

func NewSiteScrapingSettingController(cfg *domain.Config, logger *logrus.Logger, sqlClient *ent.Client) SiteScrapingSettingController {
	return &siteScrapingSettingController{
		cfg:                           cfg,
		logger:                        logger,
		sqlClient:                     sqlClient,
		siteScrapingSettingInteractor: private.NewSiteScrapingSettingInteractor(sqlClient),
	}
}

func (sssc *siteScrapingSettingController) GetBySiteSlug(c *gin.Context) {
	siteSlug := c.Param("slug")
	scrapingSettingModel, err := sssc.sqlClient.ScrapingSetting.Query().
		Where(scrapingsetting.HasSiteWith(site.SlugEQ(siteSlug))).
		WithSite().
		Only(c.Request.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			rest.RespondMessage(c, http.StatusNotFound, err.Error())
			return
		} else {
			sssc.logger.WithFields(logrus.Fields{
				"error": err.Error(),
				"site":  siteSlug,
			}).Error("failed to get site scraping settings")
			rest.RespondMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	data := domain.SiteScrapingSettingOutputScrapingSetting{
		SiteSlug:            scrapingSettingModel.Edges.Site.Slug,
		Selector:            scrapingSettingModel.Selector,
		InnerSelector:       scrapingSettingModel.InnerSelector,
		TitleSelector:       scrapingSettingModel.TitleSelector,
		DescriptionSelector: scrapingSettingModel.DescriptionSelector,
		LinkSelector:        &scrapingSettingModel.LinkSelector,
		CreatedAt:           scrapingSettingModel.CreatedAt.Format(time.RFC3339),
		UpdatedAt:           scrapingSettingModel.UpdatedAt.Format(time.RFC3339),
	}
	rest.RespondOKWithData(c, &data)
}

func (fc *siteScrapingSettingController) Upsert(c *gin.Context) {
	site_slug := c.Param("slug")
	siteModel, err := fc.sqlClient.Site.Query().
		Where(site.SlugEQ(site_slug)).
		WithFeeds(func(fq *ent.FeedQuery) {
			fq.WithFeedItems(func(fiq *ent.FeedItemQuery) {
				fiq.Order(feeditem.ByPublishedAt(sql.OrderDesc()))
			})
		}).
		WithScrapingSettings().
		Only(c.Request.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			rest.RespondMessage(c, http.StatusNotFound, err.Error())
			return
		}
	}

	var scrapingSetting domain.ScrapingSetting
	err = c.Bind(&scrapingSetting)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	siteScraper := scraper.NewScraper(fc.cfg, fc.logger)
	newFeed, err := siteScraper.FetchFeedElements(siteModel.URL, siteModel.EnableJsRendering, &scrapingSetting)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	var settingID int
	if err = db.WithTx(c.Request.Context(), fc.sqlClient, func(tx *ent.Tx) error {
		if len(siteModel.Edges.Feeds) == 0 {
			// create
			fc.logger.WithFields(logrus.Fields{
				"site": site_slug,
			}).Info("creating new feed")

			feedModel, err := tx.Feed.Create().
				SetSite(siteModel).
				SetTitle(newFeed.Title).
				SetDescription(newFeed.Description).
				SetLink(newFeed.Link).
				SetPublishedAt(newFeed.PublishedAt).
				Save(c.Request.Context())
			if err != nil {
				return err
			}
			err = tx.FeedItem.MapCreateBulk(newFeed.Items, func(fic *ent.FeedItemCreate, i int) {
				fic.SetFeed(feedModel).
					SetTitle(newFeed.Items[i].Title).
					SetDescription(newFeed.Items[i].Description).
					SetLink(*newFeed.Items[i].Link).
					SetPublishedAt(newFeed.Items[i].PublishedAt)
			}).Exec(c.Request.Context())
			if err != nil {
				return err
			}
		} else {
			// update
			fc.logger.WithFields(logrus.Fields{
				"site": site_slug,
			}).Info("updating existing feed")

			oldFeed := domain.ConvertFeedFromModelToDomain(siteModel.Edges.Feeds[0])

			rssUtil := rss.NewRssUtil(fc.cfg, fc.logger)
			updatedFeed := rssUtil.Update(oldFeed, newFeed)

			feedModel := siteModel.Edges.Feeds[0]
			feedModel, err = tx.Feed.UpdateOne(feedModel).
				SetTitle(updatedFeed.Title).
				SetDescription(updatedFeed.Description).
				SetLink(updatedFeed.Link).
				SetPublishedAt(updatedFeed.PublishedAt).
				Save(c.Request.Context())
			if err != nil {
				return err
			}
			_, err = tx.FeedItem.Delete().
				Where(feeditem.HasFeedWith(feed.IDEQ(feedModel.ID))).
				Exec(c.Request.Context())
			if err != nil {
				return err
			}
			err = tx.FeedItem.MapCreateBulk(updatedFeed.Items, func(fic *ent.FeedItemCreate, i int) {
				fic.SetFeed(feedModel).
					SetTitle(updatedFeed.Items[i].Title).
					SetDescription(updatedFeed.Items[i].Description).
					SetLink(*updatedFeed.Items[i].Link).
					SetPublishedAt(updatedFeed.Items[i].PublishedAt)
			}).Exec(c.Request.Context())
			if err != nil {
				return err
			}
		}

		settingID, err = tx.ScrapingSetting.Create().
			SetSite(siteModel).
			SetSelector(scrapingSetting.Selector).
			SetInnerSelector(scrapingSetting.InnerSelector).
			SetTitleSelector(scrapingSetting.TitleSelector).
			SetDescriptionSelector(scrapingSetting.DescriptionSelector).
			SetNillableLinkSelector(scrapingSetting.LinkSelector).
			OnConflict().
			UpdateNewValues().
			ID(c.Request.Context())
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		fc.logger.WithFields(logrus.Fields{
			"error": err.Error(),
			"site":  siteModel.Slug,
		}).Error("feed upsert failed")
		rest.RespondMessage(c, http.StatusInternalServerError, err.Error())
	}

	scrapingSettingModel, err := fc.sqlClient.ScrapingSetting.Query().
		Where(scrapingsetting.IDEQ(settingID)).
		WithSite().
		Only(c.Request.Context())

	if err != nil {
		rest.RespondMessage(c, http.StatusNotFound, err.Error())
		return
	}

	fc.logger.WithFields(logrus.Fields{
		"site": siteModel.Slug,
	}).Info("upserted feed")

	data := domain.SiteScrapingSettingOutputScrapingSetting{
		SiteSlug:            scrapingSettingModel.Edges.Site.Slug,
		Selector:            scrapingSettingModel.Selector,
		InnerSelector:       scrapingSettingModel.InnerSelector,
		TitleSelector:       scrapingSettingModel.TitleSelector,
		DescriptionSelector: scrapingSettingModel.DescriptionSelector,
		LinkSelector:        &scrapingSettingModel.LinkSelector,
		CreatedAt:           scrapingSettingModel.CreatedAt.Format(time.RFC3339),
		UpdatedAt:           scrapingSettingModel.UpdatedAt.Format(time.RFC3339),
	}
	rest.RespondOKWithData(c, &data)
}

func (fc *siteScrapingSettingController) Delete(c *gin.Context) {
	site_slug := c.Param("slug")
	siteModel, err := fc.sqlClient.Site.Query().
		Where(site.SlugEQ(site_slug), site.HasFeeds()).
		WithFeeds().
		WithScrapingSettings().
		Only(c.Request.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			rest.RespondMessage(c, http.StatusNotFound, err.Error())
			return
		}
	}

	if err = db.WithTx(c.Request.Context(), fc.sqlClient, func(tx *ent.Tx) error {
		err = tx.Feed.DeleteOne(siteModel.Edges.Feeds[0]).Exec(c.Request.Context())
		if err != nil {
			return err
		}
		err = tx.ScrapingSetting.DeleteOne(siteModel.Edges.ScrapingSettings[0]).Exec(c.Request.Context())
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		fc.logger.WithFields(logrus.Fields{
			"error": err.Error(),
			"site":  site_slug,
		}).Error("failed to delete feed")
		rest.RespondMessage(c, http.StatusInternalServerError, err.Error())
	}

	rest.RespondOKWithData(c, gin.H{
		"feed_id": siteModel.Edges.Feeds[0].ID.String(),
		"site":    site_slug,
	})
}
