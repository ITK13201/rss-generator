package private

import (
	"context"

	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/ent/site"
)

type SiteScrapingSettingInteractor interface {
	GetSite(siteID int) (*ent.Site, error)
}

type siteScrapingSettingInteractor struct {
	sqlClient *ent.Client
}

func NewSiteScrapingSettingInteractor(sqlClient *ent.Client) SiteScrapingSettingInteractor {
	return &siteScrapingSettingInteractor{
		sqlClient: sqlClient,
	}
}

func (fi *siteScrapingSettingInteractor) GetSite(siteID int) (*ent.Site, error) {
	ctx := context.Background()
	s, err := fi.sqlClient.Site.Query().Where(site.IDEQ(siteID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return s, nil
}
