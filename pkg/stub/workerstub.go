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

type WorkerStub struct {
	Name       string
	TaskQueue  string
	Workflows  []Task
	Activities []Task
	worker     worker.Worker
}

func (workerStub WorkerStub) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Starting worker '%s'. Task queue: '%s'\n", workerStub.Name, workerStub.TaskQueue)

	c, err := client.Dial(client.Options{})
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
