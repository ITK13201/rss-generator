package controllers

import (
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/interfaces/interactors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type FeedController interface {
	Get(ctx *gin.Context)
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

func (fc *feedController) Get(c *gin.Context) {
}
