package main

import (
	"gorten/internal/gorten/api"
	"gorten/internal/gorten/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	r := gin.Default()

	api.SetupRoutes(r)

	r.Run() // Listen to 0.0.0.0:8080
}
