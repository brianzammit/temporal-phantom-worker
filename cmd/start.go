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
			console.Error("Error loading config")
			os.Exit(1)
		}

		var wg sync.WaitGroup

		// TODO: Handle cleanup
		for _, workerConfig := range config.Workers {

			workerStub := stub.WorkerStub{
				Name:       workerConfig.Name,
				TaskQueue:  workerConfig.TaskQueue,
				Workflows:  workflowStubsFromConfig(workerConfig.Workflows),
				Activities: activityStubsFromConfig(workerConfig.Activities),
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

func activityStubsFromConfig(activitiesConfig []configuration.Activity) []stub.Activity {
	stubs := make([]stub.Activity, len(activitiesConfig))
	for i, c := range activitiesConfig {
		stubs[i] = stub.Activity{
			Type:   c.Type,
			Result: c.Result,
		}
	}
	return stubs
}

func workflowStubsFromConfig(workflowsConfig []configuration.Workflow) []stub.Workflow {
	stubs := make([]stub.Workflow, len(workflowsConfig))
	for i, c := range workflowsConfig {
		stubs[i] = stub.Workflow{
			Type:   c.Type,
			Result: c.Result,
		}
	}
	return stubs
}
