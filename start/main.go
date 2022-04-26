package main

import (
	"context"
	"fmt"
	"log"

	"go.temporal.io/sdk/client"

	"hello-world-project-template-go/app"
)

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	//Start workflow for "Hello String!"
	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: app.GreetingTaskQueue,
	}
	name := "Priyal"
	we, err := c.ExecuteWorkflow(context.Background(), options, app.GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}
	var greeting string
	err = we.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}
	printResults(greeting, we.GetID(), we.GetRunID())

	//Start workflow for InsertOneDocument

	options1 := client.StartWorkflowOptions{
		ID:        "InsertDoc-workflow",
		TaskQueue: app.InsertDocTaskQueue,
	}

	we1, err := c.ExecuteWorkflow(context.Background(), options1, app.MongoConnectionWorkflow)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}
	var greeting1 string
	err = we1.Get(context.Background(), &greeting1)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}
	printResults(greeting1, we1.GetID(), we1.GetRunID())

}

func printResults(greeting string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", greeting)
}
