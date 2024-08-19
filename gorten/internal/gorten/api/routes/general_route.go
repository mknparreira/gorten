package routes

import (
	"gorten/internal/gorten/api/handlers"

	"github.com/gin-gonic/gin"
)

func GeneralRoute(r *gin.Engine) *gin.RouterGroup {
	route := r.Group("/api/v1")
	{
		route.GET("/ping", handlers.PingHandler)
	}
	return route
}
