package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//wrapper for mongo client
type mongoLogWriter struct {
	client *mongo.Client
	db     *mongo.Database
}

//it disconnects mongo db client
func DisconnectMongoWriter(ctx context.Context, mongoWriter *mongoLogWriter) {
	//At the end of the program it is a best practice to close the connection:
	err := mongoWriter.client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

// open a connection to mongodb and returns the mongoLogWriter
func NewMongoLogWriter(ctx context.Context) (*mongoLogWriter, error) {

	//TODO: define the mongodb uri somewhere else
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB

	writer, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, errors.Wrap(err, "mongo disconnection failed")
	}
	LogDb := writer.Database("fleetLog")
	return &mongoLogWriter{client: writer, db: LogDb}, nil
}

// Test the connection with mongo db
func TestMongoConnection(ctx context.Context, mongoWriter *mongoLogWriter) {
	err := mongoWriter.client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to MongoDB!")
	}
}

func (mw *mongoLogWriter) Write(ctx context.Context, logs []json.RawMessage) error {
	for _, log := range logs {
		snapLog := new(Snaplog)
		b_log := []byte(log)
		fmt.Printf("this is the log casted as []byte :\n%s\n\n", b_log)
		err := json.Unmarshal(b_log, &snapLog)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("this is the result of the unmarshal:\n%#v\n\n", snapLog)
		fmt.Println("decorations: ")
		fmt.Println(snapLog.Dec)
		for i, q_res := range snapLog.Snapshot {
			fmt.Printf("row[%d]: %s\n", i, q_res)
		}

		fmt.Println("creating collection: " + snapLog.QueryName)
		collection := mw.db.Collection(snapLog.QueryName) //it creates the collection if it doesn't exist
		result, err := collection.InsertOne(ctx, snapLog)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Inserted this document: %s", result.InsertedID)
	}

	return nil
}
