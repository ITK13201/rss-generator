package rss

import (
	"github.com/ITK13201/rss-generator/domain"
	feedsLib "github.com/gorilla/feeds"
	"github.com/sirupsen/logrus"
)

type RssUtil struct {
	cfg    *domain.Config
	logger *logrus.Logger
}

func NewRssUtil(cfg *domain.Config, logger *logrus.Logger) *RssUtil {
	return &RssUtil{
		cfg:    cfg,
		logger: logger,
	}
}

func (r *RssUtil) Generate(f *domain.Feed) (*string, error) {
	feed := &feedsLib.Feed{
		Title:       f.Title,
		Description: f.Description,
		Link:        &feedsLib.Link{Href: f.Link},
		Created:     f.PublishedAt,
	}
	feedItems := []*feedsLib.Item{}
	for _, item := range f.Items {
		feedItem := &feedsLib.Item{
			Title:       item.Title,
			Description: item.Description,
			Link:        &feedsLib.Link{Href: *item.Link},
			Created:     item.PublishedAt,
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

func (r *RssUtil) Update(oldFeed *domain.Feed, newFeed *domain.Feed) *domain.Feed {
	updatedFeedItems := []*domain.FeedItem{}
	lastUpdatedAt := oldFeed.PublishedAt
	for i := 0; i < len(newFeed.Items); i++ {
		newFeedItem := newFeed.Items[i]
		updatedFeedItem := &domain.FeedItem{
			Title:       newFeedItem.Title,
			Description: newFeedItem.Description,
			Link:        newFeedItem.Link,
			PublishedAt: newFeedItem.PublishedAt,
		}
		for i := 0; i < len(oldFeed.Items); i++ {
			if oldFeed.Items[i].Title == newFeedItem.Title {
				updatedFeedItem.PublishedAt = oldFeed.Items[i].PublishedAt
				break
			}
		}
		if lastUpdatedAt.Before(updatedFeedItem.PublishedAt) {
			lastUpdatedAt = updatedFeedItem.PublishedAt
		}
		updatedFeedItems = append(updatedFeedItems, updatedFeedItem)
	}

	updatedFeed := &domain.Feed{
		Title:       newFeed.Title,
		Description: newFeed.Description,
		Link:        newFeed.Link,
		PublishedAt: lastUpdatedAt,
		Items:       updatedFeedItems,
	}
	return updatedFeed
}
