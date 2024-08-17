package db_test

import (
	"context"
	"gorten/internal/gorten/db"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMain(m *testing.M) {
	envPath := filepath.Join("../../../../", ".env")
	env := godotenv.Load(envPath)
	if env != nil {
		log.Println("Failed:", envPath)
		log.Fatalf("Failed to load .env: %v", envPath)
	}
	m.Run()
}

func setupConnection() (context.Context, context.CancelFunc) {
	ctx, cancel := db.Connect(os.Getenv("MONGODB_CONNECT_URL"))
	return ctx, cancel
}

func TestConnect(t *testing.T) {
	ctx, cancel := setupConnection()
	defer db.Disconnect(ctx, cancel)

	testCtx, testCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer testCancel()

	err := db.MongoClient.Ping(testCtx, nil)

	require.NoError(t, err, "Failed to ping MongoDB")
	assert.NotNil(t, db.MongoClient, "MongoClient should not be nil")
}

func TestDisconnect(t *testing.T) {
	// Setup a connection
	ctx, cancel := setupConnection()
	defer db.Disconnect(ctx, cancel)

	// Disconnect the client
	db.Disconnect(ctx, cancel)

	// Reinitialize context with a timeout to ensure the test doesn't hang
	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer timeoutCancel()

	// Attempt to ping the client after disconnecting
	err := db.MongoClient.Ping(timeoutCtx, nil)

	// Expect that Ping will fail after disconnecting
	require.Error(t, err, "Expected error after disconnecting")
	assert.Contains(t, err.Error(), "client is disconnected", "Expected client disconnected error, but got a different error")
}

func TestConnectTimeout(t *testing.T) {
	uri := "mongodb://10.255.255.1:27017" // IP not available

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	require.NotNil(t, client, "Expected client to be initialized")

	err = client.Ping(ctx, nil)
	if err != nil {
		t.Logf("Expected error due to timeout: %v", err)
		assert.Contains(t, err.Error(), "context deadline exceeded", "Expected context deadline exceeded error, but got a different error")
		require.Error(t, err, "Expected error due to timeout, but got none")
		return
	}
}
