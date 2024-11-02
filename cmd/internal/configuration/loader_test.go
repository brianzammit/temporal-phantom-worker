package configuration

import (
	"os"
	"testing"
)

func TestValidateAndLoad(t *testing.T) {
	tests := []struct {
		name          string
		filename      string
		expectedError bool
	}{
		{
			name:          "Valid YAML file",
			filename:      "valid_config.yaml", // Create this test file in the same directory for this test to pass
			expectedError: false,
		},
		{
			name:          "Invalid YAML file",
			filename:      "invalid_config.yaml", // Create this test file with invalid content
			expectedError: true,
		},
		{
			name:          "File does not exist",
			filename:      "non_existent.yaml",
			expectedError: true,
		},
	}

	// Create test files
	defer os.Remove("valid_config.yaml")
	defer os.Remove("invalid_config.yaml")

	// Valid file content
	validContent := `
workers:
  - name: worker1
    task_queue: queue1
    workflows:
      - type: workflow1
        result: Hello World
    activities:
      - type: activity1
        error:
          type: errorType
          message: oops
`
	// Invalid file content
	invalidContent := `invalid_yaml_content`

	// Create the valid YAML file
	if err := os.WriteFile("valid_config.yaml", []byte(validContent), 0644); err != nil {
		t.Fatalf("Failed to create valid config file: %v", err)
	}

	// Create the invalid YAML file
	if err := os.WriteFile("invalid_config.yaml", []byte(invalidContent), 0644); err != nil {
		t.Fatalf("Failed to create invalid config file: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ValidateAndLoad(tt.filename)

			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}
		})
	}
}
