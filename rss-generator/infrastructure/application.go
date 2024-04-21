package infrastructure

import (
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/interfaces/controllers"
	"github.com/sirupsen/logrus"
)

type Application struct {
	SiteController     controllers.SiteController
	TestFeedController controllers.TestFeedController
}

func NewApplication(cfg *domain.Config, logger *logrus.Logger, sqliClient *ent.Client) *Application {
	return &Application{
		SiteController:     controllers.NewSiteController(cfg, logger, sqliClient),
		TestFeedController: controllers.NewTestFeedController(cfg, logger, sqliClient),
	}
}
