package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"gorten/internal/gorten/api/providers"
	"gorten/internal/gorten/config"
)

func main() {
	app := fx.New(
		fx.Provide(
			providers.ConfigProvider,
			providers.DatabaseProvider,
			providers.MiddlewaresProvider,
		),
		providers.ModulesProvider(),
		fx.Invoke(startServer),
	)

	//It starts all registered initialization hooks, including those that are part
	//of the fx.Lifecycle and functions invoked via fx.Invoke.
	if err := app.Start(context.Background()); err != nil {
		log.Fatalf("Failed to start app: %v", err)
	}
}

func startServer(r *gin.Engine, cfg *config.AppConfig) {
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
