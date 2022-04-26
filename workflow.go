package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	var result string
	err := workflow.ExecuteActivity(ctx, ComposeGreeting, name).Get(ctx, &result)
	return result, err
}

func MongoConnectionWorkflow(ctx2 workflow.Context) (string, error) {
	//setup workflow properties
	options1 := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	//execute the activity of creating a single document
	ctx2 = workflow.WithActivityOptions(ctx2, options1)
	var result string
	err := workflow.ExecuteActivity(ctx2, MongoSingleInsert).Get(ctx2, &result)
	return result, err

}
