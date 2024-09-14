package routes

import (
	"gorten/internal/gorten/api/handlers"

	"github.com/gin-gonic/gin"
)

func CategoryRoute(r *gin.Engine, categoryHandler *handlers.CategoryHandler) *gin.RouterGroup {
	v1 := r.Group("/api/v1")
	{
		categories := v1.Group("/categories")
		{
			categories.GET("/", categoryHandler.List)
			categories.GET("/:id", categoryHandler.GetByID)
			categories.PUT("/:id", categoryHandler.UpdateByID)
			categories.POST("/", categoryHandler.Create)
		}
	}
	return v1
}
