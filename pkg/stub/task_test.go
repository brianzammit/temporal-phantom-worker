package stub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSuccessTask(t *testing.T) {
	taskType := "SuccessTask"
	taskResult := *new(interface{})
	task, _ := NewSuccessTask(taskType, taskResult)

	// Assert task properties
	assert.Equal(t, taskType, task.Type)
	assert.NotNil(t, task.result)
	assert.Nil(t, task.error)
	assert.NotNil(t, &taskResult, task.result)

	// Assert that the task is considered successful
	assert.True(t, task.isSuccess())
}

func TestNewErrorTask(t *testing.T) {
	const taskType = "ErrorTask"
	const errorType = "CustomError"
	const errorMessage = "An error occurred"
	const details = "Additional details here"
	taskError := Error{
		Type: "CustomError",
	}
	task, _ := NewErrorTask(taskType, errorType, errorMessage, details)

	// Assert task properties
	assert.Equal(t, taskType, task.Type)
	assert.Nil(t, task.result)
	assert.NotNil(t, task.error)
	assert.NotNil(t, &taskError, task.error)

	// Assert that the task is not considered successful
	assert.False(t, task.isSuccess())
}

func TestTask_isSuccess(t *testing.T) {
	// Test for a successful task
	successTask := Task{
		Type:   "Success",
		result: new(ResultTemplate), // non-nil result
		error:  nil,
	}
	assert.True(t, successTask.isSuccess())

	// Test for a failed task
	errorTask := Task{
		Type:   "Error",
		result: nil,
		error:  &Error{Type: "ErrorType", Message: new(ResultTemplate)},
	}
	assert.False(t, errorTask.isSuccess())
}
