package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

//It opens an existing collection or make a new one. The name of the collection is the name of the query
func openCollection(ctx context.Context, client *mongo.Client, log GeneralLog) {
	//TODO: Create a collection for each query name
	//the idea is to use the name of the query as identifier

}

func writeLog() {
	//TODO: write the log in the collection
	//write the log inside the collection opened with openCollection
}
