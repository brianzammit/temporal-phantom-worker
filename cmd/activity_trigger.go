package cmd

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

var triggerActivityCmd = &cobra.Command{
	Use:     "trigger",
	Aliases: []string{"t"},
	Short:   "Trigger an activity by wrapping it in a workflow",
	Long:    `Trigger an activity by wrapping it in a workflow`,
	Example: `
	# Trigger an activity through a workflow
	./phantom-worker activity trigger -type MyTestActivity -taskqueue testQueue
		`,
	Run: func(cmd *cobra.Command, args []string) {
		activityType, _ := cmd.Flags().GetString("type")
		taskQueue, _ := cmd.Flags().GetString("taskqueue")
		inputFile, _ := cmd.Flags().GetString("input")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetInt("port")
		namespace, _ := cmd.Flags().GetString("namespace")
		certPath, _ := cmd.Flags().GetString("certPath")
		keyPath, _ := cmd.Flags().GetString("keyPath")

		fmt.Printf("Triggerring activity '%s' on taskqueue '%s' - input file: '%s'\n", activityType, taskQueue, inputFile)

		var input interface{}

		if inputFile != "" {
			file, err := os.ReadFile(inputFile)
			if err != nil {
				log.Fatalln(fmt.Errorf("error reading input file: %s", err))
			}

			err = yaml.Unmarshal(file, &input)
			if err != nil {
				log.Fatalln(fmt.Errorf("error parsing input file: %s", err))
			}
		} else {
			input = nil
		}

		triggerActivity(activityType, taskQueue, input, host, port, namespace, certPath, keyPath)
	},
}

func triggerActivity(activityType string, taskQueue string, input interface{}, host string, port int, namespace string, certPath string, keyPath string) {
	id, _ := uuid.NewUUID()
	workerQueue := id.String() // Each worker is temporary and unique. Generate a unique task queue name for it
	workflowName := activityType + "TriggerWorkflow"

	c, w, err := startWorker(activityType, taskQueue, host, port, namespace, certPath, keyPath, workerQueue, workflowName)
	defer w.Stop()

	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: workerQueue,
	}

	// Now that we have the worker setup, trigger the workflow immediately
	wf, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflowName, input)
	if err != nil {
		log.Fatalln("Unable to execute workflow.", err)
	}

	var result interface{}

	err = wf.Get(context.Background(), &result)
	if err != nil {
		log.Printf("\n---\nError:\n%s\n---\n", err)
		return
	}

	yamlResult, err := yaml.Marshal(result)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Unable to marshal workflow result: %s", err))
	}

	log.Printf("\n---\nResult:\n%s---\n", yamlResult)
}

func startWorker(activityType string, taskQueue string, host string, port int, namespace string, certPath string, keyPath string, workerQueue string, workflowName string) (client.Client, worker.Worker, error) {
	clientOpts := buildClientOpts(host, port, namespace, certPath, keyPath)
	fmt.Printf("Connecting to Temporal @ '%s:%d'\n", host, port)

	// Register a workflow that triggers the desired activity
	c, err := client.Dial(clientOpts)
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}

	w := worker.New(c, workerQueue, worker.Options{
		// do not register workflows and activities by their function name
		DisableRegistrationAliasing: true,
	})

	// Register a temporary workflow that is responsible for triggering the desired activity
	w.RegisterWorkflowWithOptions(func(ctx workflow.Context, input interface{}) (interface{}, error) {
		// Set the desired task Queue
		ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			TaskQueue:              taskQueue,
			ScheduleToCloseTimeout: time.Hour,
			RetryPolicy: &temporal.RetryPolicy{
				MaximumAttempts: 1,
			},
		})

		// execute the desired activity
		var result interface{}
		err := workflow.ExecuteActivity(ctx, activityType, input).Get(ctx, &result)
		if err != nil {
			return nil, err
		}

		return result, nil
	}, workflow.RegisterOptions{
		Name: workflowName,
	})

	err = w.Start()
	if err != nil {
		log.Fatalln(fmt.Errorf("error starting worker: %s", err))
	}
	return c, w, err
}

func buildClientOpts(host string, port int, namespace string, certPath string, keyPath string) client.Options {
	clientOpts := client.Options{
		HostPort:  fmt.Sprintf("%s:%d", host, port),
		Namespace: namespace,
	}

	if certPath != "" && keyPath != "" {
		cert, err := tls.LoadX509KeyPair(certPath, keyPath)
		if err != nil {
			log.Fatalln("Failed to load X509 certificate and key. Error:", err)
		}

		clientOpts.ConnectionOptions = client.ConnectionOptions{
			TLS: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
		}
	}

	return clientOpts
}

func init() {
	triggerActivityCmd.Flags().StringP("type", "t", "", "The Activity Type")
	triggerActivityCmd.MarkFlagRequired("type")
	triggerActivityCmd.Flags().StringP("taskqueue", "q", "", "The task queue")
	triggerActivityCmd.MarkFlagRequired("taskqueue")
	triggerActivityCmd.Flags().StringP("input", "i", "", "The file containing yaml input to pass to the activity")
	triggerActivityCmd.Flags().StringP("host", "s", "localhost", "The temporal server hostname")
	triggerActivityCmd.Flags().IntP("port", "p", 7233, "The temporal server port")
	triggerActivityCmd.Flags().StringP("namespace", "n", "", "The temporal namespace")
	triggerActivityCmd.Flags().StringP("cert_path", "c", "", "The Temporal mTLS certificate path")
	triggerActivityCmd.Flags().StringP("key_path", "k", "", "The Temporal mTLS key path")

	activityCmd.AddCommand(triggerActivityCmd)
}
