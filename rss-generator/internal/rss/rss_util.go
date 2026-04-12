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
			Created:     item.PublishedAt,
		}
		if item.Link != nil {
			feedItem.Link = &feedsLib.Link{Href: *item.Link}
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
	oldFeedItemsByTitle := map[string]*domain.FeedItem{}
	for _, oldFeedItem := range oldFeed.Items {
		if oldFeedItem.Title != "" {
			if _, exists := oldFeedItemsByTitle[oldFeedItem.Title]; !exists {
				oldFeedItemsByTitle[oldFeedItem.Title] = oldFeedItem
			}
		}
	}

	for i := 0; i < len(newFeed.Items); i++ {
		newFeedItem := newFeed.Items[i]
		updatedFeedItem := &domain.FeedItem{
			Title:       newFeedItem.Title,
			Description: newFeedItem.Description,
			Link:        newFeedItem.Link,
			PublishedAt: newFeedItem.PublishedAt,
		}

		if oldFeedItem, exists := oldFeedItemsByTitle[newFeedItem.Title]; exists {
			updatedFeedItem.PublishedAt = oldFeedItem.PublishedAt
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
