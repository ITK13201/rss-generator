package interactors

import (
	"context"
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/ent/site"
	"time"
)

type TestFeedInteractor interface {
	GetSite(siteID int) (*ent.Site, error)
	ParseFeed(f *ent.Feed) *domain.Feed
	CreateFeed(siteID int, f *domain.Feed) (*string, error)
}

type testFeedInteractor struct {
	sqlClient *ent.Client
}

func NewTestFeedInteractor(sqlClient *ent.Client) TestFeedInteractor {
	return &testFeedInteractor{
		sqlClient: sqlClient,
	}
}

func (tfi *testFeedInteractor) GetSite(siteID int) (*ent.Site, error) {
	ctx := context.Background()
	s, err := tfi.sqlClient.Site.Query().Where(site.IDEQ(siteID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (tfi *testFeedInteractor) ParseFeed(f *ent.Feed) *domain.Feed {
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

func (tfi *testFeedInteractor) CreateFeed(siteID int, f *domain.Feed) (*string, error) {
	ctx := context.Background()
	now := time.Now()
	feed, err := tfi.sqlClient.Feed.Create().
		SetSiteID(siteID).
		SetTitle(f.Title).
		SetDescription(f.Description).
		SetLink(f.Link).
		SetPublishedAt(now).
		SetIsTest(true).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	_, err = tfi.sqlClient.FeedItem.MapCreateBulk(f.Items, func(fic *ent.FeedItemCreate, i int) {
		fic.SetFeed(feed).
			SetTitle(f.Items[i].Title).
			SetDescription(f.Items[i].Description).
			SetLink(*f.Items[i].Link).
			SetPublishedAt(now)
	}).Save(ctx)
	if err != nil {
		return nil, err
	}
	uuid := feed.ID.String()
	return &uuid, nil
}
