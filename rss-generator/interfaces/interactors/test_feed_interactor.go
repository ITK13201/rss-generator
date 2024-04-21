package interactors

import (
	"context"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/ent/site"
)

type TestFeedInteractor interface {
	GetSite(siteID int) (*ent.Site, error)
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
