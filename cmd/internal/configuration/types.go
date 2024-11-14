package configuration

import (
	"errors"
	"fmt"
)

type Activity struct {
	Type   string        `yaml:"type"`
	Result *interface{}  `yaml:"result"`
	Error  *ErrorDetails `yaml:"error"`
}

type Workflow struct {
	Type   string        `yaml:"type"`
	Result *interface{}  `yaml:"result"`
	Error  *ErrorDetails `yaml:"error"`
}

type ErrorDetails struct {
	Message string      `yaml:"message"`
	Type    string      `yaml:"type"`
	Details interface{} `yaml:"details"`
}

type Worker struct {
	Name       string     `yaml:"name"`
	TaskQueue  string     `yaml:"task_queue"`
	Workflows  []Workflow `yaml:"workflows"`
	Activities []Activity `yaml:"activities"`
}

type Mtls struct {
	CertPath string `yaml:"cert_path"`
	KeyPath  string `yaml:"key_path"`
}

type Server struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Namespace string `yaml:"namespace"`
	Mtls      *Mtls  `yaml:"mtls"`
}

type Config struct {
	Server  *Server  `yaml:"server"`
	Workers []Worker `yaml:"workers"`
}

func (a Activity) IsSuccess() bool {
	return a.Error == nil
}

func (w Workflow) IsSuccess() bool {
	return w.Error == nil
}

// Semantic validations
func (config Config) validate() error {
	errorMessages := make([]string, 0)
	errorMessages = append(errorMessages, config.validateUniqueWorkerNames()[:]...)
	errorMessages = append(errorMessages, config.validateUniqueTaskQueueNames()...)
	errorMessages = append(errorMessages, config.validatedNonEmptyWorkers()...)

	if len(errorMessages) > 0 {
		for _, errMessage := range errorMessages {
			fmt.Printf("Validation error: %s\n", errMessage)
		}
		return errors.New("configuration is invalid")
	} else {
		return nil
	}
}

func (config Config) validateUniqueWorkerNames() []string {
	errorMessages := make([]string, 0)

	workerNameCounts := make(map[string]int)
	for _, worker := range config.Workers {
		workerNameCounts[worker.Name] = workerNameCounts[worker.Name] + 1
	}

	for name, count := range workerNameCounts {
		if count > 1 {
			errorMessages = append(errorMessages, fmt.Sprintf("Worker '%s' configured %d times. Worker names should be unique.", name, count))
		}
	}

	return errorMessages
}

func (config Config) validateUniqueTaskQueueNames() []string {
	errorMessages := make([]string, 0)

	taskQueueNameCount := make(map[string]int)
	for _, worker := range config.Workers {
		taskQueueNameCount[worker.TaskQueue] = taskQueueNameCount[worker.TaskQueue] + 1
	}

	for name, count := range taskQueueNameCount {
		if count > 1 {
			errorMessages = append(errorMessages, fmt.Sprintf("Task Queue '%s' used by %d workers. Different worker types should use different task queues.", name, count))
		}
	}

	return errorMessages
}

func (config Config) validatedNonEmptyWorkers() []string {
	errorMessages := make([]string, 0)

	for _, worker := range config.Workers {
		if len(worker.Workflows)+len(worker.Activities) == 0 {
			errorMessages = append(errorMessages, fmt.Sprintf("Worker '%s' does not handle any Activity or Workflow. Each worker should handle at least 1 Activity or Workflow", worker.Name))
		}
	}

	return errorMessages
}
