package cmd

import (
	"github.com/spf13/cobra"
)

var JobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "Manage job scripts",
	Long:  "Manage job scripts",
}

func init() {
	rootCmd.AddCommand(JobsCmd)
}
