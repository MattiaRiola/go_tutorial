package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func closeDb(ctx context.Context, client *mongo.Client) {
	//At the end of the program it is a best practice to close the connection:
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

// open a connection to mongodb
func Connect(ctx context.Context, client *mongo.Client) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	//defer function to close the connection with the db once main finishes
	defer closeDb(ctx, client)
	if err != nil {
		log.Fatal(err)
	}
}
