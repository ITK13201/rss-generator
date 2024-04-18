package cmd

import (
	"fmt"
	"github.com/ITK13201/rss-generator/domain"
	"github.com/caarlos0/env/v11"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	config  domain.Config

	rootCmd = &cobra.Command{
		Use:   "rss-generator",
		Short: "Web application to create rss feeds by scraping websites",
		Long:  "Web application to create rss feeds by scraping websites",
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	err := initConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	err := env.Parse(&config)
	if err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}

	err = env.Parse(&config.Database)
	if err != nil {
		return fmt.Errorf("failed to parse database config: %w", err)
	}

	return err
}
