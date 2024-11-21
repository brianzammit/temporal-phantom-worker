package stub

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/temporal"
)

type activityStub struct {
	task Task
}

func (a activityStub) Execute(ctx context.Context, input interface{}) (interface{}, error) {
	info := activity.GetInfo(ctx)

	fmt.Printf("Processing activity '%s' on task queue '%s\n", info.ActivityType.Name, info.TaskQueue)
	fmt.Printf("Input: %s\n", input)

	if a.task.isSuccess() {
		processedResult, err := a.task.result.process(input)
		if err != nil {
			return nil, err
		}

		fmt.Printf("Output: %+v\n", processedResult)

		return processedResult, nil
	} else {
		fmt.Printf("Error output: %v\n", a.task.error)

		var message string
		err := processResult(*a.task.error.Message, input, &message)
		if err != nil {
			return nil, err
		}

		var details interface{}
		err = processResult(*a.task.error.Details, input, &details)
		if err != nil {
			return nil, err
		}

		return nil, temporal.NewApplicationError(
			message,
			a.task.error.Type,
			details)
	}
}

func processResult[T interface{}](resultTemplate ResultTemplate, input interface{}, output *T) error {
	details, err := resultTemplate.process(input)
	if err != nil {
		return err
	}
	*output = details.(T)
	return nil
}
