package private

import (
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
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
	var input domain.TestFeedCreateInput
	err := c.Bind(&input)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	site, err := tfc.testFeedInteractor.GetSite(input.SiteID)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	siteScraper := scraper.NewScraper(tfc.cfg, tfc.logger)
	f, err := siteScraper.FetchFeedElements(site.URL, site.EnableJsRendering, &input.Selectors)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	feedID, err := tfc.testFeedInteractor.CreateFeed(site.ID, f)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	rest.RespondOKWithData(c, gin.H{
		"feed_id": *feedID,
	})
}
