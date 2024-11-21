package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
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

		serverConfig := stub.ServerConfiguration{
			Host:      "localhost",
			Port:      7233,
			Namespace: "default",
		}
		if config.Server != nil {
			if len(config.Server.Host) > 0 {
				serverConfig.Host = config.Server.Host
			}

			if config.Server.Port != 0 {
				serverConfig.Port = config.Server.Port
			}

			if len(config.Server.Namespace) > 0 {
				serverConfig.Namespace = config.Server.Namespace
			}

			if config.Server.Mtls != nil {
				serverConfig.Mtls = &stub.MtlsConfiguration{
					CertPath: config.Server.Mtls.CertPath,
					KeyPath:  config.Server.Mtls.KeyPath,
				}
			}
		}

		// TODO: Handle cleanup
		for _, workerConfig := range config.Workers {

			workerStub := stub.WorkerStub{
				Name:         workerConfig.Name,
				TaskQueue:    workerConfig.TaskQueue,
				Workflows:    taskFromWorkflowConfig(workerConfig.Workflows),
				Activities:   taskFromActivityConfig(workerConfig.Activities),
				ServerConfig: serverConfig,
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
			task, err := stub.NewSuccessTask(a.Type, a.Result)
			if err != nil {
				log.Fatalln(fmt.Errorf("error creating success activity task: %v", err))
			}
			stubs[i] = *task
		} else {
			task, err := stub.NewErrorTask(a.Type, a.Error.Type, a.Error.Message, a.Error.Details)
			if err != nil {
				log.Fatalln(fmt.Errorf("error creating error acitivty task: %v", err))
			}
			stubs[i] = *task
		}
	}
	return stubs
}

func taskFromWorkflowConfig(workflowsConfig []configuration.Workflow) []stub.Task {
	stubs := make([]stub.Task, len(workflowsConfig))
	for i, w := range workflowsConfig {
		if w.IsSuccess() {
			task, err := stub.NewSuccessTask(w.Type, w.Result)
			if err != nil {
				log.Fatalln(fmt.Errorf("error creating success workflow task: %v", err))
			}
			stubs[i] = *task
		} else {
			task, err := stub.NewErrorTask(w.Type, w.Error.Type, w.Error.Message, w.Error.Details)
			if err != nil {
				log.Fatalln(fmt.Errorf("error creating error workflow task: %v", err))
			}
			stubs[i] = *task
		}
	}
	return stubs
}
