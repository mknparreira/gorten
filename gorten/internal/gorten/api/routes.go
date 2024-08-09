package api

import (
	"gorten/internal/gorten/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", handlers.PingHandler)
}
