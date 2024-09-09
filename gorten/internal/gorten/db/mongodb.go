package db

import (
	"context"
	"fmt"
	"gorten/pkg/logs"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClientInterface interface {
	Ping(ctx context.Context, rp *readpref.ReadPref) error
	Disconnect(ctx context.Context) error
	Database(name string, opts ...*options.DatabaseOptions) *mongo.Database
}

var MongoClient MongoClientInterface

func Connect(uri string) (context.Context, context.CancelFunc, *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//Remove resources associated with the context as soon as they are no longer needed
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		logs.Logger.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logs.Logger.Fatalf("Failed to ping MongoDB: %v", err)
	}

	MongoClient = client
	logs.Logger.Println("MongoDB connected!")

	return ctx, cancel, client
}

func Disconnect(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	if MongoClient != nil {
		if err := MongoClient.Disconnect(ctx); err != nil {
			logs.Logger.Printf("Error while disconnecting from MongoDB: %v", err)
		} else {
			fmt.Println("Closed connection")
		}
	}
}

func GetCollection(databaseName, collectionName string) *mongo.Collection {
	return MongoClient.Database(databaseName).Collection(collectionName)
}
