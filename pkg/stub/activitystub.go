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
		fmt.Printf("Output: %s\n", a.task.result)

		return a.task.result, nil
	} else {
		fmt.Printf("Error output: %s\n", a.task.error)
		return nil, temporal.NewApplicationError(
			a.task.error.Message,
			a.task.error.Type,
			a.task.error.Details)
	}
}
