package infrastructure

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

func NewRouter(publicApp *PublicApplication, privateApp *PrivateApplication) *gin.Engine {
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
	}

	// private endpoints
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			sites := v1.Group("/sites")
			{
				sites.POST("", privateApp.SiteController.Create)
			}
			testFeeds := v1.Group("/test-feeds")
			{
				testFeeds.POST("", privateApp.TestFeedController.Create)
			}
			feeds := v1.Group("/feeds")
			{
				feeds.GET("")
			}
		}
	}

	router.NoRoute()

	return router
}
