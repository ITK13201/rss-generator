package private

import (
	"fmt"
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/ent/feed"
	"github.com/ITK13201/rss-generator/ent/site"
	"github.com/ITK13201/rss-generator/interfaces/interactors/private"
	"github.com/ITK13201/rss-generator/internal/rest"
	"github.com/ITK13201/rss-generator/internal/scraper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type FeedController interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type feedController struct {
	cfg            *domain.Config
	logger         *logrus.Logger
	sqlClient      *ent.Client
	feedInteractor private.FeedInteractor
}

func NewFeedController(cfg *domain.Config, logger *logrus.Logger, sqlClient *ent.Client) FeedController {
	return &feedController{
		cfg:            cfg,
		logger:         logger,
		sqlClient:      sqlClient,
		feedInteractor: private.NewFeedInteractor(sqlClient),
	}
}

func (fc *feedController) Create(c *gin.Context) {
	var input domain.FeedCreateInput
	err := c.Bind(&input)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	s, err := fc.sqlClient.Site.Query().Where(site.IDEQ(input.SiteID)).WithFeeds().Only(c.Request.Context())
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	if len(s.Edges.Feeds) != 0 {
		rest.RespondMessage(c, http.StatusBadRequest, fmt.Sprintf("feed of site id '%d' already exists", input.SiteID))
		return
	}
	siteScraper := scraper.NewScraper(fc.cfg, fc.logger)
	f, err := siteScraper.FetchFeedElements(s.URL, s.EnableJsRendering, &input.Selectors)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	feedID, err := fc.feedInteractor.CreateFeed(c.Request.Context(), s.ID, f)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	rest.RespondOKWithData(c, gin.H{
		"feed_id": feedID,
	})
}
func (fc *feedController) Delete(c *gin.Context) {
	id := c.Param("id")
	idUUID, err := uuid.FromBytes([]byte(id))
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
	}
	f, err := fc.sqlClient.Feed.Query().Where(feed.IDEQ(idUUID)).Only(c.Request.Context())
	if err != nil {
		if ent.IsNotFound(err) {
			rest.RespondMessage(c, http.StatusNotFound, err.Error())
			return
		} else {
			fc.logger.WithFields(logrus.Fields{
				"error": err.Error(),
				"id":    id,
			}).Error("feed query failed")
			rest.RespondMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	err = fc.sqlClient.Feed.DeleteOne(f).Exec(c.Request.Context())
	if err != nil {
		fc.logger.WithFields(logrus.Fields{
			"error": err.Error(),
			"id":    id,
		}).Error("feed delete failed")
		rest.RespondMessage(c, http.StatusInternalServerError, err.Error())
	}
	rest.RespondOKWithData(c, gin.H{"feed_id": id})
}
