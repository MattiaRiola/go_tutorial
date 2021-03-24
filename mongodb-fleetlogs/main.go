package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("#######\tWELCOME\t#######")
	var mongoClient *mongo.Client
	mongoConnect(context.TODO(), mongoClient)

	fmt.Println("Bye bye")

}
