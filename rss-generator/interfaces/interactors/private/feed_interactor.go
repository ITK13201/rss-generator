package private

import (
	"context"
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/ent/site"
	"time"
)

type FeedInteractor interface {
	GetSite(siteID int) (*ent.Site, error)
	CreateFeed(ctx context.Context, siteID int, f *domain.Feed) (*string, error)
}

type feedInteractor struct {
	sqlClient *ent.Client
}

func NewFeedInteractor(sqlClient *ent.Client) FeedInteractor {
	return &feedInteractor{
		sqlClient: sqlClient,
	}
}

func (fi *feedInteractor) GetSite(siteID int) (*ent.Site, error) {
	ctx := context.Background()
	s, err := fi.sqlClient.Site.Query().Where(site.IDEQ(siteID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (fi *feedInteractor) CreateFeed(ctx context.Context, siteID int, f *domain.Feed) (*string, error) {
	now := time.Now()
	feed, err := fi.sqlClient.Feed.Create().
		SetSiteID(siteID).
		SetTitle(f.Title).
		SetDescription(f.Description).
		SetLink(f.Link).
		SetPublishedAt(now).
		SetIsTest(false).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	_, err = fi.sqlClient.FeedItem.MapCreateBulk(f.Items, func(fic *ent.FeedItemCreate, i int) {
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
