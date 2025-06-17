package private

import (
	"context"

	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
)

type SiteInteractor interface {
	Create(site *domain.SitesCreateInput) (*ent.Site, error)
}

type siteInteractor struct {
	sqlClient *ent.Client
}

func NewSiteInteractor(sqlClient *ent.Client) SiteInteractor {
	return &siteInteractor{
		sqlClient: sqlClient,
	}
}

func (si *siteInteractor) Create(site *domain.SitesCreateInput) (*ent.Site, error) {
	ctx := context.Background()
	s, err := si.sqlClient.Site.
		Create().
		SetSlug(site.Slug).
		SetTitle(site.Title).
		SetNillableDescription(site.Description).
		SetURL(site.URL).
		SetEnableJsRendering(site.EnableJSRendering).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return s, nil
}
