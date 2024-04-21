package rss

import (
	"github.com/ITK13201/rss-generator/domain"
	feedsLib "github.com/gorilla/feeds"
	"github.com/sirupsen/logrus"
	"time"
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

func (r *RssGenerator) Generate(f *domain.Feed) (*string, error) {
	var createdAt time.Time
	if f.PublishedAt == nil {
		createdAt = time.Now()
	} else {
		createdAt = *f.PublishedAt
	}
	feed := &feedsLib.Feed{
		Title:       f.Title,
		Description: f.Description,
		Link:        &feedsLib.Link{Href: f.Link},
		Created:     createdAt,
	}
	feedItems := []*feedsLib.Item{}
	for _, item := range f.Items {
		var itemCreatedAt time.Time
		if item.PublishedAt == nil {
			itemCreatedAt = time.Now()
		} else {
			itemCreatedAt = *item.PublishedAt
		}
		feedItem := &feedsLib.Item{
			Title:       item.Title,
			Description: item.Description,
			Link:        &feedsLib.Link{Href: *item.Link},
			Created:     itemCreatedAt,
		}
		feedItems = append(feedItems, feedItem)
	}
	feed.Items = feedItems

	rss, err := feed.ToRss()
	if err != nil {
		return nil, err
	}

	return &rss, nil
}
