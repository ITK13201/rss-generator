package public

import (
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
)

type TestFeedInteractor interface {
	ParseFeed(f *ent.TestFeed) *domain.Feed
}

type testFeedInteractor struct {
	sqlClient *ent.Client
}

func NewTestFeedInteractor(sqlClient *ent.Client) TestFeedInteractor {
	return &testFeedInteractor{
		sqlClient: sqlClient,
	}
}

func (tfi *testFeedInteractor) ParseFeed(f *ent.TestFeed) *domain.Feed {
	feed := &domain.Feed{
		Title:       f.Title,
		Description: f.Description,
		Link:        f.Link,
		PublishedAt: f.PublishedAt,
	}
	feedItems := []*domain.FeedItem{}
	for _, item := range f.Edges.TestFeedItems {
		feedItem := &domain.FeedItem{
			Title:       item.Title,
			Description: item.Description,
			Link:        &item.Link,
			PublishedAt: item.PublishedAt,
		}
		feedItems = append(feedItems, feedItem)
	}
	feed.Items = feedItems
	return feed
}
