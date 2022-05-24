package workflows

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type SecretDetails struct {
	Name string
}

func ProvisionSecretWorkflow(ctx workflow.Context, secretDetails SecretDetails) error {
	// RetryPolicy specifies how to automatically handle retries if an Activity fails.
	retrypolicy := &temporal.RetryPolicy{
	    InitialInterval:    time.Second,
	    BackoffCoefficient: 2.0,
	    MaximumInterval:    time.Minute,
	    MaximumAttempts:    500,
	}
	options := workflow.ActivityOptions{
	    // Timeout options specify when to automatically timeout Activity functions.
	    StartToCloseTimeout: time.Minute,
	    // Optionally provide a customized RetryPolicy.
	    // Temporal retries failures by default, this is just an example.
	    RetryPolicy: retrypolicy,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	// TODO Setup - provide activity
	// err := workflow.ExecuteActivity(ctx, ProvisionSecretStore, secretDetails).Get(ctx, nil)
	// if err != nil {
	//     return err
	// }
	// TODO Get Secret details and interact with the resource
	// err = workflow.ExecuteActivity(ctx, SetSecret, secretDetails).Get(ctx, nil)
	// if err != nil {
	//     return err
	// }

	// TODO Cleanup - provide activity
	// err = workflow.ExecuteActivity(ctx, DeprovisionSecretStore, secretDetails).Get(ctx, nil)
	// if err != nil {
	// 	return err
	// }
	return nil
}
