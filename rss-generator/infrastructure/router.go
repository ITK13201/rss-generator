package infrastructure

import (
	"net/http"

	"github.com/ITK13201/rss-generator/domain"
	"github.com/gin-gonic/gin"
)

func corsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Max-Age", "43200")

		// If the request method is OPTIONS, respond with 204 No Content
		if c.Request.Method == http.MethodOptions {
			c.Status(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func NewRouter(cfg *domain.Config, publicApp *PublicApplication, privateApp *PrivateApplication) *gin.Engine {
	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	// CORS middleware
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
				feeds.GET("/:id", privateApp.FeedController.GetByID)
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
