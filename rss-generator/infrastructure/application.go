package infrastructure

import (
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/interfaces/controllers/private"
	"github.com/sirupsen/logrus"
)

type Application struct {
	SiteController     private.SiteController
	TestFeedController private.TestFeedController
}

func NewApplication(cfg *domain.Config, logger *logrus.Logger, sqliClient *ent.Client) *Application {
	return &Application{
		SiteController:     private.NewSiteController(cfg, logger, sqliClient),
		TestFeedController: private.NewTestFeedController(cfg, logger, sqliClient),
	}
}
