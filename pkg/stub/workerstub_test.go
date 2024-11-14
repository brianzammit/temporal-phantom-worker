package stub

import (
	"crypto/tls"
	"github.com/stretchr/testify/assert"
	"go.temporal.io/sdk/testsuite"
	"sync"
	"testing"
)

func TestServerConfiguration_toTemporalOptions_noMtls(t *testing.T) {
	// Setup valid server configuration without mtls
	serverConfig := ServerConfiguration{
		Host:      "test.hostname.com",
		Port:      7234,
		Namespace: "my-test-namespace",
	}

	// Execute
	clientOpts := serverConfig.toTemporalOptions()

	// Validate
	assert.Equal(t, "test.hostname.com:7234", clientOpts.HostPort)
	assert.Equal(t, "my-test-namespace", clientOpts.Namespace)
	assert.Nil(t, clientOpts.ConnectionOptions.TLS)
}

func TestServerConfiguration_toTemporalOptions_withMtls(t *testing.T) {
	// Mock the Certificate Loading
	originalLoadX509KeyPair := loadX509KeyPair
	defer func() {
		loadX509KeyPair = originalLoadX509KeyPair
	}()
	loadX509KeyPair = func(certFile, keyFile string) (tls.Certificate, error) { return tls.Certificate{}, nil }

	// Setup valid server configuration with mtls
	serverConfig := ServerConfiguration{
		Host:      "test.hostname.com",
		Port:      7234,
		Namespace: "my-test-namespace",
		Mtls: &MtlsConfiguration{
			CertPath: "testdata/cert.pem",
			KeyPath:  "testdata/key.pem",
		},
	}

	// Execute
	clientOpts := serverConfig.toTemporalOptions()

	// Validate
	assert.Equal(t, "test.hostname.com:7234", clientOpts.HostPort)
	assert.Equal(t, "my-test-namespace", clientOpts.Namespace)
	assert.NotNil(t, clientOpts.ConnectionOptions.TLS)
}

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
