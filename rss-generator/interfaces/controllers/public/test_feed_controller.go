package public

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/ent"
	"github.com/ITK13201/rss-generator/ent/testfeed"
	"github.com/ITK13201/rss-generator/ent/testfeeditem"
	"github.com/ITK13201/rss-generator/interfaces/interactors/public"
	"github.com/ITK13201/rss-generator/internal/rest"
	"github.com/ITK13201/rss-generator/internal/rss"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type TestFeedController interface {
	Get(ctx *gin.Context)
}

type testFeedController struct {
	cfg                *domain.Config
	logger             *logrus.Logger
	sqlClient          *ent.Client
	testFeedInteractor public.TestFeedInteractor
}

func NewTestFeedController(cfg *domain.Config, logger *logrus.Logger, sqlClient *ent.Client) TestFeedController {
	return &testFeedController{
		cfg:                cfg,
		logger:             logger,
		sqlClient:          sqlClient,
		testFeedInteractor: public.NewTestFeedInteractor(sqlClient),
	}
}

func (tfc *testFeedController) Get(c *gin.Context) {
	ctx := context.Background()

	feedID := c.Param("id")
	feedUUID, err := uuid.Parse(feedID)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	f, err := tfc.sqlClient.TestFeed.Query().
		Where(testfeed.IDEQ(feedUUID)).
		WithSite().
		WithTestFeedItems(func(tfiq *ent.TestFeedItemQuery) {
			tfiq.Order(
				testfeeditem.ByPublishedAt(sql.OrderDesc()),
			)
		}).
		Only(ctx)
	if err != nil {
		rest.RespondMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	parsedFeed := tfc.testFeedInteractor.ParseFeed(f)
	rssUtil := rss.NewRssUtil(tfc.cfg, tfc.logger)
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
