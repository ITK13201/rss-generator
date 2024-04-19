package infrastructure

import (
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/interfaces/controllers"
	"github.com/sirupsen/logrus"
)

type Application struct {
	SiteController controllers.SiteController
}

func NewApplication(cfg *domain.Config, logger *logrus.Logger, sqliClient *ent.Client) *Application {
	siteController := controllers.NewSiteController(cfg, logger, sqliClient)
	return &Application{
		SiteController: siteController,
	}
}
