package cmd

import (
	"github.com/spf13/cobra"
)

var stubCmd = &cobra.Command{
	Use:     "stub",
	Aliases: []string{"s"},
	Short:   "Use help to see a list of subcommands",
}

func init() {
	rootCmd.AddCommand(stubCmd)
}
