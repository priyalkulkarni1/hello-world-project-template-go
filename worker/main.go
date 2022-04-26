package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"hello-world-project-template-go/app"
)

//There are multiple activities is being done by the same worker

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, app.GreetingTaskQueue, worker.Options{})
	//These calls are for "Hello String!"
	w.RegisterWorkflow(app.GreetingWorkflow)
	w.RegisterActivity(app.ComposeGreeting)
	//These calls are for "Insert Single document into MongoDB"
	w.RegisterWorkflow(app.MongoConnectionWorkflow)
	w.RegisterActivity(app.MongoSingleInsert)

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
