package api

import (
	"gorten/internal/gorten/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", handlers.PingHandler)

		users := v1.Group("/users")
		{
			users.GET("/", handlers.ListUsers)
			users.GET("/:id", handlers.ListUser)
			users.PUT("/:id", handlers.UpdateUser)
			users.POST("/", handlers.CreateUser)
		}
	}
}
