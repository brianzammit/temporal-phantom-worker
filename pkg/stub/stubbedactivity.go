package stub

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/activity"
)

type StubbedActivity struct {
	Result interface{}
}

func (a StubbedActivity) Execute(ctx context.Context, input interface{}) (interface{}, error) {
	info := activity.GetInfo(ctx)

	fmt.Printf("Processing activity '%s' on task queue '%s\n", info.ActivityType.Name, info.TaskQueue)
	fmt.Printf("Input: %s\n", input)
	fmt.Printf("Output: %s\n", a.Result)

	return a.Result, nil
}
