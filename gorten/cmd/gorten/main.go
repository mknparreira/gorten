package main

import (
	"github.com/joho/godotenv"

	"gorten/internal/gorten/api"
	"gorten/internal/gorten/api/middlewares"
	"gorten/internal/gorten/config"
	"gorten/internal/gorten/db"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	env := godotenv.Load()
	if env != nil {
		log.Fatalf("Failed to load .env: %v", env)
	}

	ctx, cancel := db.Connect(os.Getenv("MONGODB_CONNECT_URL"))
	//Ensure disconnect after execution
	defer db.Disconnect(ctx, cancel)

	r := gin.Default()
	r.Use(middlewares.ErrorHandlerMiddleware())

	api.SetupRoutes(r)

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
