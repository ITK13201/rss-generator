package private

import (
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/interfaces/interactors/private"
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

func (fc *feedController) Get(c *gin.Context) {
}
