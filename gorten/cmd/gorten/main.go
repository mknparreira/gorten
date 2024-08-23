package main

import (
	"context"
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/fx"

	"gorten/internal/gorten/api/providers"
)

func main() {
	envPath := filepath.Join("../../", ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Failed to load .env: %v", err)
	}

	app := fx.New(
		fx.Provide(
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

func startServer(r *gin.Engine) {
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
