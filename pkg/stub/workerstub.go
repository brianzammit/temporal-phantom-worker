package stub

import (
	"crypto/tls"
	"fmt"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"log"
	"sync"
)

type ServerConfiguration struct {
	Host      string
	Port      int
	Namespace string
	Mtls      *MtlsConfiguration
}

type MtlsConfiguration struct {
	CertPath string
	KeyPath  string
}

type WorkerStub struct {
	Name         string
	TaskQueue    string
	Workflows    []Task
	Activities   []Task
	ServerConfig ServerConfiguration
	worker       worker.Worker
}

func (serverConfig ServerConfiguration) toTemporalOptions() client.Options {
	clientOpts := client.Options{
		HostPort:  fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port),
		Namespace: serverConfig.Namespace,
	}

	if serverConfig.Mtls != nil {
		cert, err := tls.LoadX509KeyPair(serverConfig.Mtls.CertPath, serverConfig.Mtls.KeyPath)
		if err != nil {
			log.Fatalln("Failed to load X509 certificate and key. Error: %v", err)
		}

		clientOpts.ConnectionOptions = client.ConnectionOptions{
			TLS: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
		}
	}

	return clientOpts
}

func (workerStub WorkerStub) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Starting worker '%s'. Task queue: '%s'. Server: '%s@%s:%d'\n", workerStub.Name, workerStub.TaskQueue,
		workerStub.ServerConfig.Namespace, workerStub.ServerConfig.Host, workerStub.ServerConfig.Port)

	c, err := client.Dial(workerStub.ServerConfig.toTemporalOptions())
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}

	workerStub.worker = worker.New(c, workerStub.TaskQueue, worker.Options{
		// do not register workflows and activities by their function name
		DisableRegistrationAliasing: true,
	})

	// register all the workflow stubs
	for _, wf := range workerStub.Workflows {
		workflowStub := WorkflowStub{
			task: wf,
		}
		workerStub.worker.RegisterWorkflowWithOptions(workflowStub.Execute, workflow.RegisterOptions{
			Name: wf.Type,
		})

		fmt.Printf("Registered workflow '%s'\n", wf.Type)
	}

	// register all the activity stubs
	for _, a := range workerStub.Activities {
		activityStub := activityStub{
			task: a,
		}

		workerStub.worker.RegisterActivityWithOptions(activityStub.Execute, activity.RegisterOptions{
			Name: a.Type,
		})

		fmt.Printf("Registered activity '%s'\n", a.Type)
	}

	err = workerStub.worker.Run(worker.InterruptCh())
	if err != nil {
		fmt.Printf("error starting worker '%s': %v", workerStub.Name, err)
		return
	}
}

func (workerStub WorkerStub) Stop() {
	fmt.Printf("Stopping worker %s\n", workerStub.Name)
	workerStub.worker.Stop()
}
