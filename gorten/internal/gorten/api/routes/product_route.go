package routes

import (
	"gorten/internal/gorten/api/handlers"

	"github.com/gin-gonic/gin"
)

func ProductRoute(r *gin.Engine, productHandler *handlers.ProductHandler) *gin.RouterGroup {
	v1 := r.Group("/api/v1")
	{
		products := v1.Group("/products")
		{
			products.GET("/", productHandler.List)
			products.GET("/:id", productHandler.GetByID)
			products.PUT("/:id", productHandler.UpdateByID)
			products.POST("/", productHandler.Create)
		}
	}
	return v1
}
