package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func Connect(uri string) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//Remove resources associated with the context as soon as they are no longer needed
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	MongoClient = client
	log.Println("MongoDB connected!")

	return ctx, cancel
}

func Disconnect(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	if MongoClient != nil {
		if err := MongoClient.Disconnect(ctx); err != nil {
			log.Printf("Error while disconnecting from MongoDB: %v", err)
		} else {
			fmt.Println("Closed connection")
		}
	}
}

func GetCollection(databaseName, collectionName string) *mongo.Collection {
	return MongoClient.Database(databaseName).Collection(collectionName)
}
