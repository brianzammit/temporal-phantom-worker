package configuration

import (
	"testing"
)

func TestActivityIsSuccess(t *testing.T) {
	tests := []struct {
		name     string
		activity Activity
		expected bool
	}{
		{
			name: "Success Activity",
			activity: Activity{
				Type:   "test",
				Result: new(interface{}), // Pointer to an empty interface for success
				Error:  nil,              // No error
			},
			expected: true,
		},
		{
			name: "Error Activity",
			activity: Activity{
				Type:   "test",
				Result: nil, // No result
				Error: &ErrorDetails{
					Message: "An error occurred",
					Type:    "ErrorType",
					Details: nil,
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.activity.IsSuccess()
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestWorkflowIsSuccess(t *testing.T) {
	tests := []struct {
		name     string
		workflow Workflow
		expected bool
	}{
		{
			name: "Success Workflow",
			workflow: Workflow{
				Type:   "test",
				Result: new(interface{}), // Pointer to an empty interface for success
				Error:  nil,              // No error
			},
			expected: true,
		},
		{
			name: "Error Workflow",
			workflow: Workflow{
				Type:   "test",
				Result: nil, // No result
				Error: &ErrorDetails{
					Message: "An error occurred",
					Type:    "ErrorType",
					Details: nil,
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.workflow.IsSuccess()
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
