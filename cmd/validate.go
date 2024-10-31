package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"temporal-phantom-worker/cmd/internal/configuration"
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

		if err := configuration.ValidateYAMLFile(configFile); err != nil {
			log.Fatalf("Validation failed: %v\n", err)
		} else {
			fmt.Println("Configuration file is valid.")
		}
	},
}

func init() {
	validateCmd.Flags().StringP("config", "c", "", "Path to YAML configuration file")
	validateCmd.MarkFlagRequired("config")

	rootCmd.AddCommand(validateCmd)
}
