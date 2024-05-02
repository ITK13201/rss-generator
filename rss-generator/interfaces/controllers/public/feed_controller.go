package public

import (
	"entgo.io/ent/dialect/sql"
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/ent/feed"
	"github.com/ITK13201/rss-generator/ent/feeditem"
	"github.com/ITK13201/rss-generator/interfaces/interactors/public"
	"github.com/ITK13201/rss-generator/internal/rest"
	"github.com/ITK13201/rss-generator/internal/rss"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type FeedController interface {
	GetByID(c *gin.Context)
}

type feedController struct {
	cfg            *domain.Config
	logger         *logrus.Logger
	sqlClient      *ent.Client
	FeedInteractor public.FeedInteractor
}

func NewFeedController(cfg *domain.Config, logger *logrus.Logger, sqlClient *ent.Client) FeedController {
	return &feedController{
		cfg:            cfg,
		logger:         logger,
		sqlClient:      sqlClient,
		FeedInteractor: public.NewFeedInteractor(sqlClient),
	}
}

func (fc *feedController) GetByID(c *gin.Context) {
	feedID := c.Param("id")
	feedUUID, err := uuid.Parse(feedID)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	f, err := fc.sqlClient.Feed.Query().
		Where(feed.IDEQ(feedUUID)).
		WithSite().
		WithFeedItems(func(fiq *ent.FeedItemQuery) {
			fiq.Order(
				feeditem.ByPublishedAt(sql.OrderDesc()),
			)
		}).
		Only(c.Request.Context())
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	parsedFeed := fc.FeedInteractor.ParseFeed(f)
	rssUtil := rss.NewRssUtil(fc.cfg, fc.logger)
	rssXML, err := rssUtil.Generate(parsedFeed)
	if err != nil {
		rest.RespondMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Content-Type", "text/xml;charset=UTF-8")
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.String(http.StatusOK, *rssXML)
}
