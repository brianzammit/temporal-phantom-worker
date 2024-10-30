package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
	"temporal-phantom-worker/pkg/stub"
)

type Activity struct {
	Type   string      `json:"type"`
	Result interface{} `json:"result"`
}

type Workflow struct {
	Type   string      `json:"type"`
	Result interface{} `json:"result"`
}

type Worker struct {
	Name       string     `json:"name"`
	TaskQueue  string     `yaml:"task_queue"`
	Workflows  []Workflow `yaml:"workflows"`
	Activities []Activity `yaml:"activities"`
}

type TemporalServer struct {
	// todo: mtls config
	Target    string `yaml:"target"`
	Namespace string `yaml:"namespace"`
}

type Config struct {
	Server  TemporalServer `yaml:"server"`
	Workers []Worker       `yaml:"workers"`
}

// LoadConfig reads and parses the YAML configuration
func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func init() {
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start the worker with a YAML configuration",
		Run: func(cmd *cobra.Command, args []string) {
			configFile, _ := cmd.Flags().GetString("config")

			config, err := loadConfig(configFile)
			if err != nil {
				fmt.Println("Error loading config:", err)
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

	startCmd.Flags().StringP("config", "c", "config.yaml", "Path to YAML configuration file")

	rootCmd.AddCommand(startCmd)
}

func activityStubsFromConfig(activitiesConfig []Activity) []stub.Activity {
	stubs := make([]stub.Activity, len(activitiesConfig))
	for i, c := range activitiesConfig {
		stubs[i] = stub.Activity{
			Type:   c.Type,
			Result: c.Result,
		}
	}
	return stubs
}

func workflowStubsFromConfig(workflowsConfig []Workflow) []stub.Workflow {
	stubs := make([]stub.Workflow, len(workflowsConfig))
	for i, c := range workflowsConfig {
		stubs[i] = stub.Workflow{
			Type:   c.Type,
			Result: c.Result,
		}
	}
	return stubs
}
