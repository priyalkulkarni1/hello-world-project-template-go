package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	//Step 1: Write "Hello World!"
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 40,
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	ctx1 := workflow.WithActivityOptions(ctx, options)
	logger := workflow.GetLogger(ctx)
	var result string
	err := workflow.ExecuteActivity(ctx1, ComposeGreeting, name).Get(ctx1, &result)
	if err != nil {
		logger.Info("Workflow completed with greeting writing failed", "Error", err)
		return "", err
	}
	//return result, err

	//Step 2: Execute creation of Single Document
	ctx2 := workflow.WithActivityOptions(ctx, options)
	var result1 string
	err1 := workflow.ExecuteActivity(ctx2, MongoSingleInsert).Get(ctx2, &result1)
	if err1 != nil {
		logger.Info("Workflow completed with payment failed.", "Error", err1)
		return "", err1
	}
	logger.Info("Workflow completed with printing greeting and writing document.")
	return "COMPLETED", nil
}
