package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

const (
	MONGOURI = "mongodb://localhost:27017"
)

func main() {
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("#######\tWELCOME\t#######")

	// -) Creation of mongo writer (establish the connection with mongoDB):

	mongoWriter, err := NewMongoLogWriter(MONGOURI, context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	defer DisconnectMongoWriter(context.TODO(), mongoWriter)
	TestMongoConnection(context.TODO(), mongoWriter)

	// -) Logs Management:

	//Exercises and example:
	// randomExample()

	var log0 = json.RawMessage(`{"snapshot":[{"mtime":"128","path":"/Users/mattia/go/src/github.com/MattiaRiola/"},{"mtime":"96","path":"/Users/mattia/go/src/github.com/gorilla/"},{"mtime":"96","path":"/Users/mattia/go/src/github.com/julienschmidt/"}],"action":"snapshot","name":"pack/myPack1/mattia paths","hostIdentifier":"be45d5e9-cf57-4d9d-a9ff-38ca8fb369f9","calendarTime":"Tue Mar 30 08:11:50 2021 UTC","unixTime":"1617091910","epoch":0,"counter":0,"numerics":false,"decorations":{"host_uuid":"C3B43FB2-5E30-523C-AA51-7EE128D6B0B3","hostname":"servizi-mbp16-m.local"}}`)
	var log1 = json.RawMessage(`{"snapshot":[{"count(*)":"10"}],"action":"snapshot","name":"pack/myPack1/chrome processes","hostIdentifier":"be45d5e9-cf57-4d9d-a9ff-38ca8fb369f9","calendarTime":"Tue Mar 30 08:11:50 2021 UTC","unixTime":"1617091910","epoch":0,"counter":0,"numerics":false,"decorations":{"host_uuid":"C3B43FB2-5E30-523C-AA51-7EE128D6B0B3","hostname":"servizi-mbp16-m.local"}}`)
	var log2 = json.RawMessage(`{"snapshot":[{"mtime":"128","path":"/Users/mattia/go/src/github.com/MattiaRiola/"},{"mtime":"96","path":"/Users/mattia/go/src/github.com/gorilla/"},{"mtime":"96","path":"/Users/mattia/go/src/github.com/julienschmidt/"}],"action":"snapshot","name":"pack/myPack1/mattia paths","hostIdentifier":"be45d5e9-cf57-4d9d-a9ff-38ca8fb369f9","calendarTime":"Tue Mar 31 09:20:53 2021 UTC","unixTime":"1617091910","epoch":0,"counter":0,"numerics":false,"decorations":{"host_uuid":"C3B43FB2-5E30-523C-AA51-7EE128D6B0B3","hostname":"servizi-mbp16-m.local"}}`)
	var logs []json.RawMessage = []json.RawMessage{log0, log1, log2}
	err = mongoWriter.Write(context.TODO(), logs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bye bye")

}
