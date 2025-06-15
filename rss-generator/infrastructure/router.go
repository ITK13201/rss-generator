package infrastructure

import (
	"net/http"
	"time"

	"github.com/ITK13201/rss-generator/domain"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func corsHandler() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func NewRouter(cfg *domain.Config, publicApp *PublicApplication, privateApp *PrivateApplication) *gin.Engine {
	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Use(corsHandler())

	// ping: always return 200 OK
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// public endpoints
	{
		testFeeds := router.Group("/test-feeds")
		{
			testFeeds.GET("/:id", publicApp.TestFeedController.Get)
		}
		feeds := router.Group("/feeds")
		{
			feeds.GET("/:id", publicApp.FeedController.GetByID)
		}
	}

	// private endpoints
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			sites := v1.Group("/sites")
			{
				sites.GET("", privateApp.SiteController.GetAll)
				sites.GET("/:slug", privateApp.SiteController.GetBySlug)
				sites.POST("", privateApp.SiteController.Create)
				sites.DELETE("/:slug", privateApp.SiteController.Delete)
				sites.POST("/:slug", privateApp.SiteController.Update)

				settings := sites.Group("/:slug/settings")
				{
					settings.GET("", privateApp.SiteScrapingSettingController.GetBySiteSlug)
					settings.POST("", privateApp.SiteScrapingSettingController.Upsert)
					settings.DELETE("", privateApp.SiteScrapingSettingController.Delete)
				}
			}
			feeds := v1.Group("/feeds")
			{
				feeds.GET("", privateApp.FeedController.GetAll)
			}
			testFeeds := v1.Group("/test-feeds")
			{
				testFeeds.POST(":site_slug", privateApp.TestFeedController.Create)
			}
		}
	}

	router.NoRoute()

	return router
}
