package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Version    = "dev" // Injected during build
	BuildTime  = "n/a" // Injected during build
	CommitHash = "n/a" // Injected during build
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Print the version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version:    %s\n", Version)
		fmt.Printf("Build Time: %s\n", BuildTime)
		fmt.Printf("Commit:     %s\n", CommitHash)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
