package main

import (
	"gorten/internal/gorten/api"
	"gorten/internal/gorten/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	r := gin.Default()

	api.SetupRoutes(r)

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
