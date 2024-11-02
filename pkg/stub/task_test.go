package stub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSuccessTask(t *testing.T) {
	taskType := "SuccessTask"
	taskResult := *new(interface{})
	task := NewSuccessTask(taskType, taskResult)

	// Assert task properties
	assert.Equal(t, taskType, task.Type)
	assert.NotNil(t, task.result)
	assert.Nil(t, task.error)
	assert.Equal(t, &taskResult, task.result)

	// Assert that the task is considered successful
	assert.True(t, task.isSuccess())
}

func TestNewErrorTask(t *testing.T) {
	taskType := "ErrorTask"
	taskError := Error{
		Type:    "CustomError",
		Message: "An error occurred",
		Details: "Additional details here",
	}
	task := NewErrorTask(taskType, taskError)

	// Assert task properties
	assert.Equal(t, taskType, task.Type)
	assert.Nil(t, task.result)
	assert.NotNil(t, task.error)
	assert.Equal(t, &taskError, task.error)

	// Assert that the task is not considered successful
	assert.False(t, task.isSuccess())
}

func TestTask_isSuccess(t *testing.T) {
	// Test for a successful task
	successTask := Task{
		Type:   "Success",
		result: new(interface{}), // non-nil result
		error:  nil,
	}
	assert.True(t, successTask.isSuccess())

	// Test for a failed task
	errorTask := Task{
		Type:   "Error",
		result: nil,
		error:  &Error{Type: "ErrorType", Message: "Failure Message"},
	}
	assert.False(t, errorTask.isSuccess())
}
