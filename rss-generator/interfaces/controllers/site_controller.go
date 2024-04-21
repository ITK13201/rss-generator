package controllers

import (
	"encoding/json"
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/interfaces/interactors"
	"github.com/ITK13201/rss-generator/internal/rest"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type SiteController interface {
	Create(ctx *gin.Context)
}

type siteController struct {
	cfg            *domain.Config
	logger         *logrus.Logger
	sqlClient      *ent.Client
	siteInteractor interactors.SiteInteractor
}

func NewSiteController(cfg *domain.Config, logger *logrus.Logger, sqlClient *ent.Client) SiteController {
	return &siteController{
		cfg:            cfg,
		logger:         logger,
		sqlClient:      sqlClient,
		siteInteractor: interactors.NewSiteInteractor(sqlClient),
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
