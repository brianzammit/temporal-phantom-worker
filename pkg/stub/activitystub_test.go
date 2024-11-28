package stub

import (
	"go.temporal.io/sdk/temporal"
	"testing"

	"go.temporal.io/sdk/testsuite"
)

func TestActivityStub_Execute_Success(t *testing.T) {
	// Create a test suite
	suite := testsuite.WorkflowTestSuite{}
	env := suite.NewTestActivityEnvironment()

	// Set up a successful task
	task, _ := NewSuccessTask("exampleTask", "Success result")
	activityStub := activityStub{task: *task}

	// Define the activity execution
	env.RegisterActivity(activityStub.Execute)
	encodedValue, _ := env.ExecuteActivity(activityStub.Execute, nil)

	result := ""
	encodedValue.Get(&result)
	// Assert that the activity completes successfully
	if result != "Success result" {
		t.Errorf("Expected 'Success result', got '%s'", result)
	}
}

func TestActivityStub_Execute_Error(t *testing.T) {
	// Create a test suite
	suite := testsuite.WorkflowTestSuite{}
	env := suite.NewTestActivityEnvironment()

	const errorType = "java.io.IOException"
	const errorMessage = "oops - something went wrong"
	const errorDetail = "My detail string"

	// Set up a failing task
	task, _ := NewErrorTask("exampleTask", errorType, errorMessage, errorDetail)
	activityStub := activityStub{task: *task}

	// Define the activity execution
	env.RegisterActivity(activityStub.Execute)
	_, err := env.ExecuteActivity(activityStub.Execute, nil)

	// Assert that the activity fails with the expected error
	if err == nil {
		t.Error("Expected an error but got none")
	} else {
		// Check if the error matches the expected error
		if actError, ok := err.(*temporal.ActivityError); ok {
			appError := actError.Unwrap().(*temporal.ApplicationError)
			if appError.Message() != errorMessage || appError.Type() != errorType {
				t.Errorf("Expected error message '%s' of type '%s', got '%s' of type '%s'",
					errorMessage, errorType, appError.Message(), appError.Type())
			}
		} else {
			t.Errorf("Expected an ApplicationError but got %v", err)
		}
	}
}
