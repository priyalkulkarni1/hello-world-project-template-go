# Multi-Activity Temporal Workflow

## Project Description
This project explores a multi-activity temporal workflow. 
The following three actions are performed:

1) A "Hello World/Hello String" statement is printed. 
This was the starting point of understanding the functioning of the four main components of a Temporal Workflow. This simple demo has been made available by Temporal and can be found at [Temporal Go Hello World](https://github.com/temporalio/samples-go/tree/main/helloworld)

2) A sleep for 30 seconds is setup. This is a simulation for a long running activity. Two important parts of long activites are correctly setting up the [Start-To-Close Timeout](https://docs.temporal.io/docs/concepts/what-is-a-start-to-close-timeout) and [Activity Heartbeat](https://docs.temporal.io/docs/go/how-to-heartbeat-an-activity-in-go)

3) An Insert Operation is performed to MongoDB collection.
A new database called "QuickStart" is created. Then a collection called "Podcasts" is created. Finally, InsertOne operation is used to push demo data into a document.

## Steps to run Program
1) You need a Temporal service running. I am using [Docker-Compose](https://docs.temporal.io/docs/clusters/quick-install/#docker-compose) to run my App locally.
   When the Temporal Cluster is running, the Temporal Web UI becomes available in your [Browser](http://localhost:8088/namespaces/default/workflows?range=last-30-days&status=ALL)
   
2) You need a MongoDB Connection setup locally. I am using [MongoDB Compass](https://www.mongodb.com/products/compass) to run my App locally.

3) Run the following command to start the worker
```
go run hello-world-project-template-go/worker/main.go
```
4) Run the following command to start the example
```
go run hello-world-project-template-go/starter/main.go
```
## Failure Simulation
Two common ways to introduce errors and test Temporal's behaviour are as follows.
### Server Crash
Unlike many modern applications that require complex leader election processes and external databases to handle failure, Temporal automatically preserves the state of your Workflow even if the server is down. You can easily test this by following these steps (again, make sure your Worker is stopped so your Workflow doesn't finish):

1) Start the Workflow again.

2) Verify the Workflow is running in the UI.

3) Shut down the Temporal server by either using 'Ctrl c' or via the Docker dashboard.

4) After the Temporal server has stopped, restart it and visit the UI.

5) Your Workflow is still there!

### Activity Error
We simulate a bug in the Activity function. Please check [Activity.go](https://github.com/priyalkulkarni1/hello-world-project-template-go/blob/master/activity.go). Comments include simulated error.
Save your changes and run the Worker. You will see the Worker complete the ComposeGreeting() Activity function, but error when it attempts the MongoSingleInsert() Activity function. The important thing to note here is that the Worker keeps retrying the MongoSingleInsert() function.
