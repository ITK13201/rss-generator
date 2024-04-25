package public

import (
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
)

type FeedInteractor interface {
	ParseFeed(f *ent.Feed) *domain.Feed
}

type feedInteractor struct {
	sqlClient *ent.Client
}

func NewFeedInteractor(sqlClient *ent.Client) FeedInteractor {
	return &feedInteractor{
		sqlClient: sqlClient,
	}
}

func (fi *feedInteractor) ParseFeed(f *ent.Feed) *domain.Feed {
	feed := &domain.Feed{
		Title:       f.Title,
		Description: f.Description,
		Link:        f.Link,
		PublishedAt: &f.PublishedAt,
	}
	feedItems := []*domain.FeedItem{}
	for _, item := range f.Edges.FeedItems {
		feedItem := &domain.FeedItem{
			Title:       item.Title,
			Description: item.Description,
			Link:        &item.Link,
			PublishedAt: &item.PublishedAt,
		}
		feedItems = append(feedItems, feedItem)
	}
	feed.Items = feedItems
	return feed
}
