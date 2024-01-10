package main

import (
	"context"
	"fmt"
	"log"
	//"time"

	"go.temporal.io/sdk/client"

	"github.com/MatthiasScholz/temporal-tfsecret/internal/pkg/workflows"
)

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{
		Namespace: "default",
		HostPort:  "127.0.0.1:7233",
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "tfsecret-test-workflow",
		TaskQueue: workflows.TFSecretTaskQueue,
		// FIXME Field does not exist anymore: WorkerStopTimeout: 30 * time.Second,
	}

	secret := workflows.SecretDetailsInput{"test_secret", "eu-central-1"}
	we, err := c.ExecuteWorkflow(context.Background(), options, workflows.ProvisionSecretWorkflow, secret)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}

	printResults(result, we.GetID(), we.GetRunID())
}

func printResults(result string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", result)
}
