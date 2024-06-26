package private

import (
	"context"
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/ent/site"
	"time"
)

type TestFeedInteractor interface {
	GetSite(siteID int) (*ent.Site, error)
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

func (tfi *testFeedInteractor) CreateFeed(siteID int, f *domain.Feed) (*string, error) {
	ctx := context.Background()
	now := time.Now()
	feed, err := tfi.sqlClient.TestFeed.Create().
		SetSiteID(siteID).
		SetTitle(f.Title).
		SetDescription(f.Description).
		SetLink(f.Link).
		SetPublishedAt(now).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	_, err = tfi.sqlClient.TestFeedItem.MapCreateBulk(f.Items, func(fic *ent.TestFeedItemCreate, i int) {
		fic.SetTestFeed(feed).
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
