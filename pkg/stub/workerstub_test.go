package stub

import (
	"go.temporal.io/sdk/testsuite"
	"sync"
	"testing"
)

func TestWorkerStub_Run(t *testing.T) {
	// Create a test suite
	suite := testsuite.WorkflowTestSuite{}
	env := suite.NewTestWorkflowEnvironment()

	// Prepare mock tasks for workflows and activities
	workflows := []Task{
		NewSuccessTask("TestWorkflow", "Workflow result"),
	}
	activities := []Task{
		NewSuccessTask("TestActivity", "Activity result"),
	}

	// Initialize the WorkerStub
	workerStub := WorkerStub{
		Name:       "TestWorker",
		TaskQueue:  "TestQueue",
		Workflows:  workflows,
		Activities: activities,
	}

	var wg sync.WaitGroup
	wg.Add(1)

	// Run the worker in a separate goroutine
	go workerStub.Run(&wg)

	// Execute the workflow in the test environment
	//env.RegisterWorkflow(WorkflowStub{task: workflows[0]}.Execute)
	//env.RegisterActivity(activityStub{task: activities[0]}.Execute)

	env.ExecuteWorkflow(WorkflowStub{task: workflows[0]}.Execute, "Test input")

	// Assert that the workflow completes successfully
	env.AssertExpectations(t)
	result := ""
	env.GetWorkflowResult(&result)
	if result != "Workflow result" {
		t.Errorf("Expected 'Workflow result', got '%s'", result)
	}
}
