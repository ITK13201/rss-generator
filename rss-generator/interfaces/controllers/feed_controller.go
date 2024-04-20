package controllers

import (
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/interfaces/interactors"
	"github.com/ITK13201/rss-generator/internal/rest"
	"github.com/ITK13201/rss-generator/internal/scraper"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type FeedController interface {
	Create(ctx *gin.Context)
}

type feedController struct {
	cfg            *domain.Config
	logger         *logrus.Logger
	sqlClient      *ent.Client
	feedInteractor interactors.FeedInteractor
}

func NewFeedController(cfg *domain.Config, logger *logrus.Logger, sqlClient *ent.Client) FeedController {
	return &feedController{
		cfg:            cfg,
		logger:         logger,
		sqlClient:      sqlClient,
		feedInteractor: interactors.NewFeedInteractor(sqlClient),
	}
}

func (fc *feedController) Create(c *gin.Context) {
	var f domain.FeedCreateInput
	err := c.Bind(&f)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	site, err := fc.feedInteractor.GetSite(f.SiteID)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
	}
	rssUtil := scraper.NewUtil(fc.cfg, fc.logger)
	_, err = rssUtil.FetchFeedElements(site.URL, site.EnableJsRendering)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
	}

	rest.RespondOK(c)
	return
}
