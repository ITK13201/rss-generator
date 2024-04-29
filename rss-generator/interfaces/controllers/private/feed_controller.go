package private

import (
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/ent/feed"
	"github.com/ITK13201/rss-generator/ent/site"
	"github.com/ITK13201/rss-generator/interfaces/interactors/private"
	"github.com/ITK13201/rss-generator/internal/db"
	"github.com/ITK13201/rss-generator/internal/rest"
	"github.com/ITK13201/rss-generator/internal/scraper"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type FeedController interface {
	Upsert(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type feedController struct {
	cfg            *domain.Config
	logger         *logrus.Logger
	sqlClient      *ent.Client
	feedInteractor private.FeedInteractor
}

func NewFeedController(cfg *domain.Config, logger *logrus.Logger, sqlClient *ent.Client) FeedController {
	return &feedController{
		cfg:            cfg,
		logger:         logger,
		sqlClient:      sqlClient,
		feedInteractor: private.NewFeedInteractor(sqlClient),
	}
}

func (fc *feedController) Upsert(c *gin.Context) {
	site_slug := c.Param("site_slug")
	siteModel, err := fc.sqlClient.Site.Query().
		Where(site.SlugEQ(site_slug)).
		WithFeeds(func(fq *ent.FeedQuery) {
			fq.Where(feed.IsTestEQ(false))
		}).
		WithScrapingSettings().
		Only(c.Request.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			rest.RespondMessage(c, http.StatusNotFound, err.Error())
			return
		}
	}

	if len(siteModel.Edges.Feeds) == 1 {
		// delete old feed
		err = fc.sqlClient.Feed.DeleteOne(siteModel.Edges.Feeds[0]).Exec(c.Request.Context())
		if err != nil {
			fc.logger.WithFields(logrus.Fields{
				"error": err.Error(),
				"site":  site_slug,
			}).Error("feed delete failed")
			rest.RespondMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	var input domain.FeedUpsertInput
	err = c.Bind(&input)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	siteScraper := scraper.NewScraper(fc.cfg, fc.logger)
	f, err := siteScraper.FetchFeedElements(siteModel.URL, siteModel.EnableJsRendering, &input.ScrapingSetting)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	feedID, err := fc.feedInteractor.CreateFeed(c.Request.Context(), siteModel.ID, f)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	err = fc.sqlClient.ScrapingSetting.Create().
		SetSite(siteModel).
		SetSelector(input.ScrapingSetting.Selector).
		SetInnerSelector(input.ScrapingSetting.InnerSelector).
		SetTitleSelector(input.ScrapingSetting.TitleSelector).
		SetDescriptionSelector(input.ScrapingSetting.DescriptionSelector).
		SetNillableLinkSelector(input.ScrapingSetting.LinkSelector).
		OnConflict().
		UpdateNewValues().
		Exec(c.Request.Context())
	if err != nil {
		fc.logger.WithFields(logrus.Fields{
			"error": err.Error(),
			"site":  site_slug,
		})
		rest.RespondMessage(c, http.StatusInternalServerError, err.Error())
	}
	rest.RespondOKWithData(c, gin.H{
		"feed_id": feedID,
	})
}
func (fc *feedController) Delete(c *gin.Context) {
	site_slug := c.Param("site_slug")
	siteModel, err := fc.sqlClient.Site.Query().
		Where(site.SlugEQ(site_slug)).
		WithFeeds(func(fq *ent.FeedQuery) {
			fq.Where(feed.IsTestEQ(false))
		}).
		WithScrapingSettings().
		Only(c.Request.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			rest.RespondMessage(c, http.StatusNotFound, err.Error())
			return
		}
	}
	if len(siteModel.Edges.Feeds) != 1 {
		rest.RespondMessage(c, http.StatusNotFound, "feed not found")
		return
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
