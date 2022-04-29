# Multi-Activity Temporal Workflow

## Project Description
This project explores a multi-activity temporal workflow. 
The following two actions are performed:

1)A "Hello World/Hello String" statement is printed. 
This was the starting point of understanding the functioning of the four main components of a Temporal Workflow. This simple demo has been made available by Temporal and can be found at [Temporal Go Hello World](https://github.com/temporalio/samples-go/tree/main/helloworld)

2)An Insert Operation is performed to MongoDB collection.
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
