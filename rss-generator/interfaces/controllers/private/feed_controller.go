package private

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/ent/feed"
	"github.com/ITK13201/rss-generator/internal/rest"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type FeedController interface {
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
}

type feedController struct {
	cfg       *domain.Config
	logger    *logrus.Logger
	sqlClient *ent.Client
}

func NewFeedController(cfg *domain.Config, logger *logrus.Logger, sqlClient *ent.Client) FeedController {
	return &feedController{
		cfg:       cfg,
		logger:    logger,
		sqlClient: sqlClient,
	}
}
func (fc *feedController) GetAll(c *gin.Context) {
	feeds, err := fc.sqlClient.Feed.Query().WithSite().All(c.Request.Context())
	if err != nil {
		fc.logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("feed query failed")
		rest.RespondMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	data := make([]domain.FeedV2GetAllOutput, len(feeds))
	for i, f := range feeds {
		data[i] = domain.FeedV2GetAllOutput{
			ID:          f.ID.String(),
			SiteSlug:    f.Edges.Site.Slug,
			Title:       f.Title,
			Description: f.Description,
			Link:        f.Link,
			CreatedAt:   f.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   f.UpdatedAt.Format(time.RFC3339),
		}
	}
	dataJson, _ := json.Marshal(data)
	fc.logger.WithFields(logrus.Fields{
		"data": string(dataJson),
	}).Info("feed data retrieved successfully")

	rest.RespondOKWithData(c, &data)
}

func (fc *feedController) GetByID(c *gin.Context) {
	id := c.Param("id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		fc.logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("invalid UUID format")
		rest.RespondMessage(c, http.StatusBadRequest, "invalid id format")
		return
	}
	feed, err := fc.sqlClient.Feed.
		Query().
		Where(feed.ID(uuidID)).
		WithSite().
		Only(c.Request.Context())
	if err != nil {
		fc.logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("feed query failed")
		rest.RespondMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	data := domain.FeedV2GetAllOutput{
		ID:          feed.ID.String(),
		SiteSlug:    feed.Edges.Site.Slug,
		Title:       feed.Title,
		Description: feed.Description,
		Link:        feed.Link,
		CreatedAt:   feed.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   feed.UpdatedAt.Format(time.RFC3339),
	}
	dataJson, _ := json.Marshal(data)
	fc.logger.WithFields(logrus.Fields{
		"data": string(dataJson),
	}).Info("feed data retrieved successfully")

	rest.RespondOKWithData(c, &data)
}
