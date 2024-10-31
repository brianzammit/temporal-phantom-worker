package configuration

import (
	"embed"
	"errors"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
	"os"
)

//go:embed assets/config-schema.json
var schemaFile embed.FS

// ValidateYAMLFile validates a YAML file against the JSON schema.
func ValidateYAMLFile(filename string) error {
	// Load the JSON schema from embedded assets
	schemaJSON, err := schemaFile.ReadFile("assets/config-schema.json")
	if err != nil {
		return fmt.Errorf("failed to load schema: %w", err)
	}

	// Load YAML file content
	yamlData, err := loadYAMLFile(filename)
	if err != nil {
		return fmt.Errorf("failed to load configuration file: %w", err)
	}

	// Create loaders for schema and YAML data
	schemaLoader := gojsonschema.NewBytesLoader(schemaJSON)
	documentLoader := gojsonschema.NewGoLoader(yamlData)

	// Perform configuration
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return fmt.Errorf("schema configuration error: %w", err)
	}

	// Collect any configuration errors
	if !result.Valid() {
		for _, desc := range result.Errors() {
			fmt.Printf("Validation error: %s\n", desc)
		}
		return errors.New("configuration invalid")
	}

	return nil
}

// loadYAMLFile reads and unmarshal a YAML file into a map.
func loadYAMLFile(filename string) (map[string]interface{}, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var yamlData map[string]interface{}
	err = yaml.Unmarshal(file, &yamlData)
	if err != nil {
		return nil, err
	}

	return yamlData, nil
}
