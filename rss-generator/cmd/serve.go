package cmd

import (
	"context"
	"fmt"
	"github.com/ITK13201/rss-generator/domain"
	"github.com/ITK13201/rss-generator/infrastructure"
	"github.com/ITK13201/rss-generator/services"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run api server",
	Long:  "Run api server",
	Run: func(cmd *cobra.Command, args []string) {
		serve(&config)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve(cfg *domain.Config) {
	logger := services.NewLogger(cfg)

	sqlClient := infrastructure.NewSqlClient(cfg)
	application := infrastructure.NewApplication(cfg, logger, sqlClient)
	router := infrastructure.NewRouter(application)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}
	logger.Info("Server exiting")
}
