package routes

import (
	"gorten/internal/gorten/api/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine, userHandler *handlers.UserHandler) *gin.RouterGroup {
	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/", userHandler.List)
			users.GET("/:id", userHandler.UserByID)
			users.PUT("/:id", userHandler.UpdateByID)
			users.POST("/", userHandler.Create)
		}
	}
	return v1
}
