package cmd

import (
	"github.com/spf13/cobra"
)

var activityCmd = &cobra.Command{
	Use:     "activity",
	Aliases: []string{"a"},
	Short:   "Use help to see a list of subcommands",
}

func init() {
	rootCmd.AddCommand(activityCmd)
}
