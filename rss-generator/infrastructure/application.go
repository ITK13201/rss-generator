package infrastructure

import (
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/interfaces/controllers/private"
	"github.com/ITK13201/rss-generator/interfaces/controllers/public"
	"github.com/sirupsen/logrus"
)

type PrivateApplication struct {
	SiteController     private.SiteController
	TestFeedController private.TestFeedController
	FeedController     private.FeedController
}

type PublicApplication struct {
	TestFeedController public.TestFeedController
	FeedController     public.FeedController
}

func NewPrivateApplication(cfg *domain.Config, logger *logrus.Logger, sqliClient *ent.Client) *PrivateApplication {
	return &PrivateApplication{
		SiteController:     private.NewSiteController(cfg, logger, sqliClient),
		TestFeedController: private.NewTestFeedController(cfg, logger, sqliClient),
		FeedController:     private.NewFeedController(cfg, logger, sqliClient),
	}
}

func NewPublicApplication(cfg *domain.Config, logger *logrus.Logger, sqliClient *ent.Client) *PublicApplication {
	return &PublicApplication{
		TestFeedController: public.NewTestFeedController(cfg, logger, sqliClient),
		FeedController:     public.NewFeedController(cfg, logger, sqliClient),
	}
}
