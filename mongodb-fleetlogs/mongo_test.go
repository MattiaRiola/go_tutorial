package main

import (
	"context"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	MONGOURI_TEST        = "mongodb://localhost:27017"
	MONGO_DB_RESULT_TEST = "fleetLogResultTest"
	MONGO_DB_STATUS_TEST = "fleetLogStatusTest"
	MONGO_TEST           = true
)

var testFunctions = [...]func(*testing.T){
	testMyAlwaysOkFunc,
	// testMyAlwaysFailFunc,
	testTestMongoConnection,
}

func FunctionName(f interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	elements := strings.Split(fullName, ".")
	return elements[len(elements)-1]
}
func TestAll(t *testing.T) {
	if !MONGO_TEST {
		t.Skip("Mongo tests not requested")
	}
	for _, f := range testFunctions {
		t.Run(FunctionName(f), f)
	}
	t.Run(FunctionName(testMyAlwaysFailFunc), testMyAlwaysFailFunc)
}
func TestAllOk(t *testing.T) {
	if !MONGO_TEST {
		t.Skip("Mongo tests not requested")
	}
	for _, f := range testFunctions {
		t.Run(FunctionName(f), f)
	}
}
func testTestMongoConnection(t *testing.T) {
	_, err := NewMongoLogWriter("", context.TODO())
	require.NotNil(t, err)

	mongoW, err := NewMongoLogWriter(MONGOURI_TEST, context.TODO())
	require.Nil(t, err)

	assert.Equal(t, mongoW.db.Name(), "fleetLog")

}

func testMyAlwaysOkFunc(t *testing.T) {
	fmt.Println("Testing the always ok func")
	assert.True(t, true)
}
func testMyAlwaysFailFunc(t *testing.T) {
	fmt.Println("Testing the always fail func")
	assert.True(t, false)
}
