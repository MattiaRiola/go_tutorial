package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

//It closes the connection of the client
func closeDb(client *mongo.Client) {
	//At the end of the program it is a best practice to close the connection:
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

//delete everything from that collection
func deleteEverything(collection *mongo.Collection) {
	result, err := collection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("number of document deleted: ", result.DeletedCount)
}

// inserting item into the collection passing the object
func insertTrainer(collection *mongo.Collection, trainer Trainer) {
	insertResult, err := collection.InsertOne(context.TODO(), trainer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

}
func playingWithMongo() {

	fmt.Printf("\n\n#####\tmongo db tutorial\t#####\n\n")
	// Creating the option for the connection:
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	//defer function to close the connection with the db once main finishes
	defer closeDb(client)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to MongoDB!")
	}
	// Now I start to use the db:

	//the collection variable will handle the Collection that is in the db
	collection := client.Database("test").Collection("trainers")

	// 	fmt.Println("--- Using BSON object ---")
	// 	fmt.Println("BSON is the binary representation of json in mongodb")
	// bson.D{{
	//     "name",
	//     bson.D{{
	//         "$in",
	//         bson.A{"Alice", "Bob"}
	//     }}
	// }}
	fmt.Println("--- CRUD Operations ---")
	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 10, "Cerculean City"}
	// brock := Trainer{"Brock", 15, "Pewter City"}

	insertTrainer(collection, ash)
	insertTrainer(collection, misty)

	//manual insert with bson:
	result1, err1 := collection.InsertOne(context.TODO(), bson.M{
		"name": "Burzi",
		"age":  7,
		"city": "Turin"})
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println("Inserted a single document: ", result1.InsertedID)

	result2, err2 := collection.DeleteMany(context.TODO(), bson.M{
		"city": "Turin",
	})
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("Deleted this document: ", result2.DeletedCount)

	//TODO : Contninue the tutorial: https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
	deleteEverything(collection)
}
