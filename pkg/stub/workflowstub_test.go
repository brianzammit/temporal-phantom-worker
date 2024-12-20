package stub

import (
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/testsuite"
	"testing"
)

func TestWorkflowStub_Execute_Success(t *testing.T) {
	// Create a test suite
	suite := testsuite.WorkflowTestSuite{}
	env := suite.NewTestWorkflowEnvironment()

	// Set up a successful task
	task, _ := NewSuccessTask("exampleTask", "Success result")
	workflowStub := WorkflowStub{task: *task}

	// Define the workflow execution
	env.ExecuteWorkflow(workflowStub.Execute, nil)

	// Assert that the workflow completes successfully
	env.AssertExpectations(t)
	result := ""
	env.GetWorkflowResult(&result)
	if result != "Success result" {
		t.Errorf("Expected 'Success result', got '%s'", result)
	}
}

func TestWorkflowStub_Execute_Error(t *testing.T) {
	// Create a test suite
	suite := testsuite.WorkflowTestSuite{}
	env := suite.NewTestWorkflowEnvironment()

	var errorType = "java.io.IOException"
	var errorMessage = "oops - something went wrong"

	// Set up a failing task
	task, _ := NewErrorTask("exampleTask", errorType, errorMessage, nil)
	workflowStub := WorkflowStub{task: *task}

	// Define the workflow execution
	env.ExecuteWorkflow(workflowStub.Execute, nil)

	// Assert that the workflow fails with the expected error
	env.AssertExpectations(t)
	err := env.GetWorkflowError()
	if err == nil {
		t.Error("Expected an error but got none")
	} else {
		// Check if the error matches the expected error
		if execError, ok := err.(*temporal.WorkflowExecutionError); ok {
			appError := execError.Unwrap().(*temporal.ApplicationError)
			if appError.Message() != errorMessage || appError.Type() != errorType {
				t.Errorf("Expected error message '%s' of type '%s', got '%s' of type '%s'",
					errorMessage, errorType, appError.Message(), appError.Type())
			}
		} else {
			t.Errorf("Expected an ApplicationError but got %v", err)
		}
	}
}
