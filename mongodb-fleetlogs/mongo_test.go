package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	MONGOURI_TEST        = "mongodb://localhost:27017"
	MONGO_DB_RESULT_TEST = "fleetLogResultTest"
	MONGO_DB_STATUS_TEST = "fleetLogStatusTest"
)

func MongoConnectionTest(t *testing.T) {
	mongoW, err := NewMongoLogWriter("", context.TODO())
	require.NotNil(t, err)

	mongoW, err = NewMongoLogWriter(MONGOURI_TEST, context.TODO())
	require.Nil(t, err)

	assert.Equal(t, mongoW.db.Name(), "fleetLog")

}
