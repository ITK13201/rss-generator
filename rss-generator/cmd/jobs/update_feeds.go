package jobs

import (
	"context"
	"github.com/ITK13201/rss-generator/infrastructure"
	"github.com/ITK13201/rss-generator/interfaces/jobs"
	"github.com/ITK13201/rss-generator/services"
	"github.com/spf13/cobra"

	"github.com/ITK13201/rss-generator/cmd"
)

var updateFeedsCmd = &cobra.Command{
	Use:   "update_feeds",
	Short: "Update feeds",
	Long:  "Update feeds by scraping sites",
	Run: func(cobraCmd *cobra.Command, args []string) {
		cfg := &cmd.Config
		logger := services.NewLogger(cfg)
		sqlClient := infrastructure.NewSqlClient(cfg)
		job := jobs.NewUpdateFeedsJob(cfg, logger, sqlClient)
		ctx := context.Background()
		job.Run(ctx)
	},
}

func init() {
	cmd.JobsCmd.AddCommand(updateFeedsCmd)
}
