package private

import (
	"encoding/json"
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/ent/site"
	"github.com/ITK13201/rss-generator/interfaces/interactors/private"
	"github.com/ITK13201/rss-generator/internal/rest"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type SiteController interface {
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
		Slug:  site.Slug,
		Title: site.Title,
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

	err = sc.sqlClient.Site.UpdateOne(s).
		SetNillableTitle(input.Title).
		SetNillableDescription(input.Description).
		SetNillableURL(input.URL).
		SetNillableEnableJsRendering(input.EnableJSRendering).
		Exec(c.Request.Context())
	if err != nil {
		sc.logger.WithFields(logrus.Fields{
			"error": err.Error(),
			"slug":  slug,
		}).Error("site update failed")
		rest.RespondMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	rest.RespondOKWithData(c, gin.H{"slug": slug})
}
