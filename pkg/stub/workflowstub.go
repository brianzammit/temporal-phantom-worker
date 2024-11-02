package stub

import (
	"fmt"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type WorkflowStub struct {
	task Task
}

func (wf WorkflowStub) Execute(ctx workflow.Context, input interface{}) (interface{}, error) {
	info := workflow.GetInfo(ctx)

	fmt.Printf("Processing workflow '%s' on task queue '%s'\n", info.WorkflowType.Name, info.TaskQueueName)
	fmt.Printf("Input: %s\n", input)

	if wf.task.isSuccess() {
		fmt.Printf("Output: %s\n", wf.task.result)
		return wf.task.result, nil
	} else {
		fmt.Printf("Error output: %s\n", wf.task.error)
		return nil, temporal.NewApplicationError(
			wf.task.error.Message,
			wf.task.error.Type,
			wf.task.error.Details)
	}
}
