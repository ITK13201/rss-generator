package interactors

import (
	"context"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/ent/site"
)

type FeedInteractor interface {
	GetSite(siteID int) (*ent.Site, error)
}

type feedInteractor struct {
	sqlClient *ent.Client
}

func NewFeedInteractor(sqlClient *ent.Client) FeedInteractor {
	return &feedInteractor{
		sqlClient: sqlClient,
	}
}

func (tfi *feedInteractor) GetSite(siteID int) (*ent.Site, error) {
	ctx := context.Background()
	s, err := tfi.sqlClient.Site.Query().Where(site.IDEQ(siteID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return s, nil
}
