package stub

import (
	"fmt"
	"go.temporal.io/sdk/workflow"
)

type WorkflowStub struct {
	Result interface{}
}

func (wf WorkflowStub) Execute(ctx workflow.Context, input interface{}) (interface{}, error) {
	info := workflow.GetInfo(ctx)

	fmt.Printf("Processing workflow '%s' on task queue '%s'\n", info.WorkflowType.Name, info.TaskQueueName)
	fmt.Printf("Input: %s\n", input)
	fmt.Printf("Output: %s\n", wf.Result)

	return wf.Result, nil
}
