package stub

import (
	"fmt"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"log"
	"sync"
)

type Workflow struct {
	Type   string
	Result interface{}
}

type Activity struct {
	Type   string
	Result interface{}
}

type WorkerStub struct {
	Name       string
	TaskQueue  string
	Workflows  []Workflow
	Activities []Activity
	worker     worker.Worker
}

func (stub WorkerStub) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Starting worker %s\n", stub.Name)

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}

	stub.worker = worker.New(c, stub.TaskQueue, worker.Options{
		// do not register workflows and activities by their function name
		DisableRegistrationAliasing: true,
	})

	// register all the workflows
	for _, wf := range stub.Workflows {
		stubbedWf := StubbedWorkflow{
			Result: wf.Result,
		}
		stub.worker.RegisterWorkflowWithOptions(stubbedWf.Execute, workflow.RegisterOptions{
			Name: wf.Type,
		})
	}

	for _, a := range stub.Activities {
		stubbedA := StubbedActivity{
			Result: a.Result,
		}

		stub.worker.RegisterActivityWithOptions(stubbedA.Execute, activity.RegisterOptions{
			Name: a.Type,
		})
	}

	err = stub.worker.Run(worker.InterruptCh())
	if err != nil {
		fmt.Printf("Error starting worker '%s': %v", stub.Name, err)
		return
	}
}

func (stub WorkerStub) Stop() {
	fmt.Printf("Stopping worker %s\n", stub.Name)
	stub.worker.Stop()
}
