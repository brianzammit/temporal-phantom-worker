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
		var result interface{}
		err := processResultTemplateInSideEffect(ctx, input, *wf.task.result, &result)
		if err != nil {
			return nil, err
		}

		fmt.Printf("Output: %+v\n", result)

		return result, nil
	} else {
		fmt.Printf("Error output: %v\n", wf.task.error)
		var message string

		err := processResultTemplateInSideEffect(ctx, input, *wf.task.error.Message, &message)
		if err != nil {
			return nil, err
		}

		var details interface{}

		err = processResultTemplateInSideEffect(ctx, input, *wf.task.error.Details, &details)
		if err != nil {
			return nil, err
		}

		return nil, temporal.NewApplicationError(
			message,
			wf.task.error.Type,
			details)
	}
}

func processResultTemplateInSideEffect[T interface{}](ctx workflow.Context, input interface{}, resultTemplate ResultTemplate, output *T) error {
	encodedResult := workflow.SideEffect(ctx, func(ctx workflow.Context) interface{} {
		result, err := resultTemplate.process(input)
		if err != nil {
			fmt.Printf("Error processing result: %s\n", err.Error())
			return nil
		}
		return result
	})
	err := encodedResult.Get(output)
	if err != nil {
		return err
	}
	return nil
}
