package private

import (
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/ent/site"
	"github.com/ITK13201/rss-generator/interfaces/interactors/private"
	"github.com/ITK13201/rss-generator/internal/rest"
	"github.com/ITK13201/rss-generator/internal/scraper"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type TestFeedController interface {
	Create(ctx *gin.Context)
}

type testFeedController struct {
	cfg                *domain.Config
	logger             *logrus.Logger
	sqlClient          *ent.Client
	testFeedInteractor private.TestFeedInteractor
}

func NewTestFeedController(cfg *domain.Config, logger *logrus.Logger, sqlClient *ent.Client) TestFeedController {
	return &testFeedController{
		cfg:                cfg,
		logger:             logger,
		sqlClient:          sqlClient,
		testFeedInteractor: private.NewTestFeedInteractor(sqlClient),
	}
}

func (tfc *testFeedController) Create(c *gin.Context) {
	site_slug := c.Param("site_slug")
	siteModel, err := tfc.sqlClient.Site.Query().
		Where(site.SlugEQ(site_slug)).
		WithScrapingSettings().
		Only(c.Request.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			rest.RespondMessage(c, http.StatusNotFound, err.Error())
			return
		} else {
			tfc.logger.WithFields(logrus.Fields{
				"error": err,
				"site":  site_slug,
			}).Error("site query failed")
			rest.RespondMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	var input domain.TestFeedCreateInput
	err = c.Bind(&input)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	siteScraper := scraper.NewScraper(tfc.cfg, tfc.logger)
	f, err := siteScraper.FetchFeedElements(siteModel.URL, siteModel.EnableJsRendering, &input.ScrapingSetting)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	feedID, err := tfc.testFeedInteractor.CreateFeed(siteModel.ID, f)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	rest.RespondOKWithData(c, gin.H{
		"feed_id": *feedID,
	})
}
