package private

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/ent/site"
	"github.com/ITK13201/rss-generator/interfaces/interactors/private"
	"github.com/ITK13201/rss-generator/internal/rest"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SiteController interface {
	GetAll(ctx *gin.Context)
	GetBySlug(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type siteController struct {
	cfg            *domain.Config
	logger         *logrus.Logger
	sqlClient      *ent.Client
	siteInteractor private.SiteInteractor
}

func NewSiteController(cfg *domain.Config, logger *logrus.Logger, sqlClient *ent.Client) SiteController {
	return &siteController{
		cfg:            cfg,
		logger:         logger,
		sqlClient:      sqlClient,
		siteInteractor: private.NewSiteInteractor(sqlClient),
	}
}

func (sc *siteController) GetAll(c *gin.Context) {
	sites, err := sc.sqlClient.Site.Query().
		WithScrapingSettings().
		All(c.Request.Context())
	if err != nil {
		sc.logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("site query failed")
		rest.RespondMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	data := make([]domain.SitesGetAllOutput, len(sites))
	for i, s := range sites {
		var scrapingSetting *domain.ScrapingSettingV2
		if len(s.Edges.ScrapingSettings) > 0 {
			scrapingSetting = &domain.ScrapingSettingV2{
				SiteSlug:            s.Slug,
				Selector:            s.Edges.ScrapingSettings[0].Selector,
				InnerSelector:       s.Edges.ScrapingSettings[0].InnerSelector,
				TitleSelector:       s.Edges.ScrapingSettings[0].TitleSelector,
				DescriptionSelector: s.Edges.ScrapingSettings[0].DescriptionSelector,
				LinkSelector:        &s.Edges.ScrapingSettings[0].LinkSelector,
				CreatedAt:           s.Edges.ScrapingSettings[0].CreatedAt.Format(time.RFC3339),
				UpdatedAt:           s.Edges.ScrapingSettings[0].UpdatedAt.Format(time.RFC3339),
			}
		}
		data[i] = domain.SitesGetAllOutput{
			Slug:              s.Slug,
			Title:             s.Title,
			Description:       &s.Description,
			URL:               s.URL,
			EnableJSRendering: s.EnableJsRendering,
			CreatedAt:         s.CreatedAt.Format(time.RFC3339),
			UpdatedAt:         s.UpdatedAt.Format(time.RFC3339),
			ScrapingSetting:   scrapingSetting,
		}
	}
	dataJson, _ := json.Marshal(data)
	sc.logger.WithFields(logrus.Fields{
		"data": string(dataJson),
	}).Info("sites retrieved")
	rest.RespondOKWithData(c, &data)
}

func (sc *siteController) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")
	s, err := sc.sqlClient.Site.
		Query().
		Where(site.SlugEQ(slug)).
		WithScrapingSettings().
		Only(c.Request.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			rest.RespondMessage(c, http.StatusNotFound, err.Error())
			return
		} else {
			sc.logger.WithFields(logrus.Fields{
				"error": err.Error(),
				"slug":  slug,
			}).Error("site query failed")
			rest.RespondMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	var scrapingSetting *domain.ScrapingSettingV2
	if s.Edges.ScrapingSettings != nil {
		scrapingSetting = &domain.ScrapingSettingV2{
			SiteSlug:            s.Slug,
			Selector:            s.Edges.ScrapingSettings[0].Selector,
			InnerSelector:       s.Edges.ScrapingSettings[0].InnerSelector,
			TitleSelector:       s.Edges.ScrapingSettings[0].TitleSelector,
			DescriptionSelector: s.Edges.ScrapingSettings[0].DescriptionSelector,
			LinkSelector:        &s.Edges.ScrapingSettings[0].LinkSelector,
			CreatedAt:           s.Edges.ScrapingSettings[0].CreatedAt.Format(time.RFC3339),
			UpdatedAt:           s.Edges.ScrapingSettings[0].UpdatedAt.Format(time.RFC3339),
		}
	}
	data := domain.SitesGetBySlugOutput{
		Slug:              s.Slug,
		Title:             s.Title,
		Description:       &s.Description,
		URL:               s.URL,
		EnableJSRendering: s.EnableJsRendering,
		CreatedAt:         s.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         s.UpdatedAt.Format(time.RFC3339),
		ScrapingSetting:   scrapingSetting,
	}
	dataJson, _ := json.Marshal(data)
	sc.logger.WithFields(logrus.Fields{
		"data": string(dataJson),
	}).Info("site retrieved")
	rest.RespondOKWithData(c, &data)
}

func (sc *siteController) Create(c *gin.Context) {
	var s domain.SitesCreateInput
	err := c.Bind(&s)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	site, err := sc.siteInteractor.Create(&s)
	if err != nil {
		sc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("site creation failed")
		rest.RespondMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	data := domain.SitesCreateOutput{
		Slug:              site.Slug,
		Title:             site.Title,
		Description:       &site.Description,
		URL:               site.URL,
		EnableJSRendering: site.EnableJsRendering,
		CreatedAt:         site.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         site.UpdatedAt.Format(time.RFC3339),
	}
	dataJson, _ := json.Marshal(data)
	sc.logger.WithFields(logrus.Fields{
		"data": string(dataJson),
	}).Info("site created")
	rest.RespondOKWithData(c, &data)
}

func (sc *siteController) Delete(c *gin.Context) {
	slug := c.Param("slug")
	s, err := sc.sqlClient.Site.Query().Where(site.SlugEQ(slug)).Only(c.Request.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			rest.RespondMessage(c, http.StatusNotFound, err.Error())
			return
		} else {
			sc.logger.WithFields(logrus.Fields{
				"error": err.Error(),
				"slug":  slug,
			}).Error("site query failed")
			rest.RespondMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	err = sc.sqlClient.Site.DeleteOne(s).Exec(c.Request.Context())
	if err != nil {
		sc.logger.WithFields(logrus.Fields{
			"error": err.Error(),
			"slug":  slug,
		}).Error("site delete failed")
		rest.RespondMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	rest.RespondOKWithData(c, gin.H{"slug": slug})
}

func (sc *siteController) Update(c *gin.Context) {
	slug := c.Param("slug")
	s, err := sc.sqlClient.Site.Query().Where(site.SlugEQ(slug)).Only(c.Request.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			rest.RespondMessage(c, http.StatusNotFound, err.Error())
			return
		} else {
			sc.logger.WithFields(logrus.Fields{
				"error": err.Error(),
				"slug":  slug,
			}).Error("site query failed")
			rest.RespondMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	var input domain.SitesUpdateInput
	err = c.Bind(&input)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	updatedSite, err := sc.sqlClient.Site.UpdateOne(s).
		SetNillableTitle(input.Title).
		SetNillableDescription(input.Description).
		SetNillableURL(input.URL).
		SetNillableEnableJsRendering(input.EnableJSRendering).
		Save(c.Request.Context())
	if err != nil {
		sc.logger.WithFields(logrus.Fields{
			"error": err.Error(),
			"slug":  slug,
		}).Error("site update failed")
		rest.RespondMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	data := domain.SitesUpdateOutput{
		Slug:              updatedSite.Slug,
		Title:             updatedSite.Title,
		Description:       &updatedSite.Description,
		URL:               updatedSite.URL,
		EnableJSRendering: updatedSite.EnableJsRendering,
		CreatedAt:         updatedSite.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         updatedSite.UpdatedAt.Format(time.RFC3339),
	}
	rest.RespondOKWithData(c, &data)
}
