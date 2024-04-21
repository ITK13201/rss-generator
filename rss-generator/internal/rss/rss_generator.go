package rss

import (
	"github.com/ITK13201/rss-generator/domain"
	"github.com/sirupsen/logrus"
)

type RssGenerator struct {
	cfg    *domain.Config
	logger *logrus.Logger
}

func NewRssGenerator(cfg *domain.Config, logger *logrus.Logger) *RssGenerator {
	return &RssGenerator{
		cfg:    cfg,
		logger: logger,
	}
}

func (r *RssGenerator) Generate() (*string, error) {
	return nil, nil
}
