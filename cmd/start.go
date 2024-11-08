package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"sync"
	"temporal-phantom-worker/cmd/internal/configuration"
	"temporal-phantom-worker/cmd/internal/console"
	"temporal-phantom-worker/pkg/stub"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the worker with a YAML configuration",
	Example: `
	# Start the Phantom Worker with a specific config
	phantom-worker start -c ./config/sample.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		configFile, _ := cmd.Flags().GetString("config")

		config, err := configuration.ValidateAndLoad(configFile)
		if err != nil {
			console.Error("error loading config")
			os.Exit(1)
		}

		var wg sync.WaitGroup

		// TODO: Handle cleanup
		for _, workerConfig := range config.Workers {

			workerStub := stub.WorkerStub{
				Name:       workerConfig.Name,
				TaskQueue:  workerConfig.TaskQueue,
				Workflows:  taskFromWorkflowConfig(workerConfig.Workflows),
				Activities: taskFromActivityConfig(workerConfig.Activities),
			}

			wg.Add(1)
			go workerStub.Run(&wg)
		}

		wg.Wait()
	},
}

func init() {
	startCmd.Flags().StringP("config", "c", "", "Path to YAML configuration file")
	startCmd.MarkFlagRequired("config")

	rootCmd.AddCommand(startCmd)
}

func taskFromActivityConfig(activitiesConfig []configuration.Activity) []stub.Task {
	stubs := make([]stub.Task, len(activitiesConfig))
	for i, a := range activitiesConfig {
		if a.IsSuccess() {
			stubs[i] = stub.NewSuccessTask(a.Type, a.Result)
		} else {
			stubs[i] = stub.NewErrorTask(a.Type, stub.Error{
				Type:    a.Error.Type,
				Message: a.Error.Message,
				Details: a.Error.Details,
			})
		}
	}
	return stubs
}

func taskFromWorkflowConfig(workflowsConfig []configuration.Workflow) []stub.Task {
	stubs := make([]stub.Task, len(workflowsConfig))
	for i, w := range workflowsConfig {
		if w.IsSuccess() {
			stubs[i] = stub.NewSuccessTask(w.Type, w.Result)
		} else {
			stubs[i] = stub.NewErrorTask(w.Type, stub.Error{
				Type:    w.Error.Type,
				Message: w.Error.Message,
				Details: w.Error.Details,
			})
		}
	}
	return stubs
}
