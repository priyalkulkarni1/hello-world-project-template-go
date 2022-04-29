package app

//original local activity used to print "Hello String!"
//adding another local acitvity to write to a document in local instance of MongoDB
//can spawn various workers to do different CRUD operations for scalibility.
import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ComposeGreeting(name string) error {
	//greeting := fmt.Sprintf("\n\nHello %s!\n\n", name)
	fmt.Printf("\n\nHello %s!\n\n", name)

	//simulating expensive processing with a 30 second sleep
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
	time.Sleep(30 * time.Second)
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())

	// ERROR SIMULATION: Switch out return comments on the return statements to simulate an error
	//return fmt.Errorf("\n\n message display did not occur due to an issue")

	return nil
}

//connection to mongodb established and this activity will insert a single document

func MongoSingleInsert() error {

	//start the mongodb connection before MongoSingleInsert
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx1, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Connect(ctx1)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx1)

	//create database and collection
	quickstartDatabase := client.Database("quickstart")
	podcastsCollection := quickstartDatabase.Collection("podcasts")

	//adding a single document to podcast collection
	podcastResult, err := podcastsCollection.InsertOne(ctx1, bson.D{
		{Key: "title", Value: "The Polyglot Developer Podcast"},
		{Key: "author", Value: "Nic Raboy"},
		{Key: "tags", Value: bson.A{"development", "programming", "coding"}},
	})

	if err != nil {
		panic(err)
	}
	// end insertOne

	// When you run this file, it should print:
	// Document inserted with ID: ObjectID("...")
	fmt.Printf("\n\nDocument inserted with ID: %s\n\n", fmt.Sprint(podcastResult.InsertedID))

	return nil
}
