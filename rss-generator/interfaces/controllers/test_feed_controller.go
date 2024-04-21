package controllers

import (
	"context"
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/ent/feed"
	"github.com/ITK13201/rss-generator/interfaces/interactors"
	"github.com/ITK13201/rss-generator/internal/rest"
	"github.com/ITK13201/rss-generator/internal/rss"
	"github.com/ITK13201/rss-generator/internal/scraper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type TestFeedController interface {
	Create(ctx *gin.Context)
	Show(ctx *gin.Context)
}

type testFeedController struct {
	cfg                *domain.Config
	logger             *logrus.Logger
	sqlClient          *ent.Client
	testFeedInteractor interactors.TestFeedInteractor
}

func NewTestFeedController(cfg *domain.Config, logger *logrus.Logger, sqlClient *ent.Client) TestFeedController {
	return &testFeedController{
		cfg:                cfg,
		logger:             logger,
		sqlClient:          sqlClient,
		testFeedInteractor: interactors.NewTestFeedInteractor(sqlClient),
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

func (tfc *testFeedController) Show(c *gin.Context) {
	ctx := context.Background()

	feedID := c.Param("id")
	feedUUID, err := uuid.Parse(feedID)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	f, err := tfc.sqlClient.Feed.Query().Where(feed.IDEQ(feedUUID)).WithSite().WithFeedItems().Only(ctx)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	parsedFeed := tfc.testFeedInteractor.ParseFeed(f)
	rssGenerator := rss.NewRssGenerator(tfc.cfg, tfc.logger)
	rssXML, err := rssGenerator.Generate(parsedFeed)
	if err != nil {
		rest.RespondMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Content-Type", "text/xml;charset=UTF-8")
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.String(http.StatusOK, *rssXML)
}
