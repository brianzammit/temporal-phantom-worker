package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"temporal-phantom-worker/cmd/internal/configuration"
	"temporal-phantom-worker/cmd/internal/console"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate the configuration file",
	Example: `
	# Validate a configuration file against the schema
	./phantom-worker validate -c ./config/sample.yaml
		`,
	Run: func(cmd *cobra.Command, args []string) {
		configFile, _ := cmd.Flags().GetString("config")
		fmt.Printf("Validating configuration file: %s\n", configFile)

		if _, err := configuration.ValidateAndLoad(configFile); err != nil {
			console.Error("Validation failed: %v\n", err)
		} else {
			console.Success("Configuration file is valid.")
		}
	},
}

func init() {
	validateCmd.Flags().StringP("config", "c", "", "Path to YAML configuration file")
	validateCmd.MarkFlagRequired("config")

	rootCmd.AddCommand(validateCmd)
}
