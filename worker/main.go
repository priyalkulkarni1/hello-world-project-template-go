package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"hello-world-project-template-go/app"
)

//spinning up two different workers for two different tasks

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	//This worker is for "Hello String!"
	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, app.GreetingTaskQueue, worker.Options{})
	w.RegisterWorkflow(app.GreetingWorkflow)
	w.RegisterActivity(app.ComposeGreeting)
	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}

	//This worker is for "Insert Single document into MongoDB"
	// This worker hosts both Workflow and Activity functions
	w1 := worker.New(c, app.InsertDocTaskQueue, worker.Options{})
	w1.RegisterWorkflow(app.MongoConnectionWorkflow)
	w1.RegisterActivity(app.MongoSingleInsert)
	// Start listening to the Task Queue
	err = w1.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
