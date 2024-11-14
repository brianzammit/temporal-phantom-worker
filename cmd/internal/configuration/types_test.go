package configuration

import (
	"github.com/stretchr/testify/assert"
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

func Test_ValidateUniqueWorkerNames_Success(t *testing.T) {
	config := Config{
		Workers: []Worker{
			{
				Name: "WorkerOne",
			},
			{
				Name: "WorkerTwo",
			},
			{
				Name: "WorkerThree",
			},
		},
	}

	errors := config.validateUniqueWorkerNames()
	assert.Empty(t, errors)
}

func Test_ValidateUniqueWorkerNames_Invalid(t *testing.T) {
	config := Config{
		Workers: []Worker{
			{
				TaskQueue: "queue_1",
			},
			{
				TaskQueue: "queue_2",
			},
			{
				TaskQueue: "queue_1",
			},
			{
				TaskQueue: "queue_2",
			},
			{
				TaskQueue: "queue_1",
			},
			{
				TaskQueue: "queue_3",
			},
		},
	}

	errors := config.validateUniqueTaskQueueNames()
	assert.Equal(t, 2, len(errors))
	assert.Equal(t, errors[0], "Task Queue 'queue_1' used by 3 workers. Different worker types should use different task queues.")
	assert.Equal(t, errors[1], "Task Queue 'queue_2' used by 2 workers. Different worker types should use different task queues.")
}

func Test_ValidateUniqueTaskQueueNames_Success(t *testing.T) {
	config := Config{
		Workers: []Worker{
			{
				TaskQueue: "queue_1",
			},
			{
				TaskQueue: "queue_2",
			},
			{
				TaskQueue: "queue_3",
			},
		},
	}

	errors := config.validateUniqueTaskQueueNames()
	assert.Empty(t, errors)
}

func Test_ValidateUniqueTaskQueueNames_Invalid(t *testing.T) {
	config := Config{
		Workers: []Worker{
			{
				Name: "WorkerOne",
			},
			{
				Name: "WorkerTwo",
			},
			{
				Name: "WorkerOne",
			},
			{
				Name: "WorkerTwo",
			},
			{
				Name: "WorkerOne",
			},
			{
				Name: "WorkerThree",
			},
		},
	}

	errors := config.validateUniqueWorkerNames()
	assert.Equal(t, 2, len(errors))
	assert.Equal(t, errors[0], "Worker 'WorkerOne' configured 3 times. Worker names should be unique.")
	assert.Equal(t, errors[1], "Worker 'WorkerTwo' configured 2 times. Worker names should be unique.")
}

func Test_ValidateNonEmptyWorkers_Success(t *testing.T) {
	config := Config{
		Workers: []Worker{
			{
				Name: "WorkerOne",
				Activities: []Activity{
					{},
				},
			},
			{
				Name: "WorkerTwo",
				Workflows: []Workflow{
					{},
				},
			},
			{
				Name: "WorkerThree",
				Activities: []Activity{
					{},
				},
				Workflows: []Workflow{
					{},
				},
			},
		},
	}

	errors := config.validatedNonEmptyWorkers()
	assert.Empty(t, errors)
}

func Test_ValidateNonEmptyWorkers_Invalid(t *testing.T) {
	config := Config{
		Workers: []Worker{
			{
				Name:       "WorkerOne",
				Activities: []Activity{},
			},
			{
				Name:      "WorkerTwo",
				Workflows: []Workflow{},
			},
			{
				Name: "WorkerThree",
			},
		},
	}

	errors := config.validatedNonEmptyWorkers()
	assert.Equal(t, 3, len(errors))
	assert.Equal(t, errors[0], "Worker 'WorkerOne' does not handle any Activity or Workflow. Each worker should handle at least 1 Activity or Workflow")
	assert.Equal(t, errors[1], "Worker 'WorkerTwo' does not handle any Activity or Workflow. Each worker should handle at least 1 Activity or Workflow")
	assert.Equal(t, errors[2], "Worker 'WorkerThree' does not handle any Activity or Workflow. Each worker should handle at least 1 Activity or Workflow")
}
