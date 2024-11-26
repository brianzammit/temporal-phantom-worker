package cmd

import (
	"github.com/spf13/cobra"
)

var triggerCmd = &cobra.Command{
	Use:     "activity",
	Aliases: []string{"a"},
	Short:   "Trigger an action. Use help to see a list of subcommands",
}

func init() {
	rootCmd.AddCommand(triggerCmd)
}
