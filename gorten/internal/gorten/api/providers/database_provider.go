package providers

import (
	"context"
	"gorten/internal/gorten/config"
	"gorten/internal/gorten/db"

	"go.mongodb.org/mongo-driver/mongo"
)

func DatabaseProvider(cfx *config.AppConfig) (*mongo.Client, context.Context, context.CancelFunc) {
	ctx, cancel, client := db.Connect(cfx.Mongo.ConnectionURL)
	return client, ctx, cancel
}
