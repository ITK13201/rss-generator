package infrastructure

import (
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/sirupsen/logrus"
)

type Application struct {
}

func NewApplication(cfg *domain.Config, logger *logrus.Logger, sqliClient *ent.Client) *Application {
	return &Application{}
}
