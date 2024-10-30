package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "phantomworker",
	Short: "Temporal Phantom Worker - a CLI for running a configurable Temporal worker stub",
	Long:  `Temporal Phantom Worker is a CLI tool for running a stubbed Temporal worker, useful for testing workflows in environments where certain services are unavailable.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
