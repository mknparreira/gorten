package providers

import (
	"context"
	"gorten/internal/gorten/db"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func DatabaseProvider() (*mongo.Client, context.Context, context.CancelFunc) {
	ctx, cancel, client := db.Connect(os.Getenv("MONGODB_CONNECT_URL"))
	return client, ctx, cancel
}
