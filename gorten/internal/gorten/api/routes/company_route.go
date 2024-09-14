package routes

import (
	"gorten/internal/gorten/api/handlers"

	"github.com/gin-gonic/gin"
)

func CompanyRoute(r *gin.Engine, companyHandler *handlers.CompanyHandler) *gin.RouterGroup {
	v1 := r.Group("/api/v1")
	{
		companies := v1.Group("/companies")
		{
			companies.GET("/", companyHandler.List)
			companies.GET("/:id", companyHandler.GetByID)
			companies.PUT("/:id", companyHandler.UpdateByID)
			companies.POST("/", companyHandler.Create)
		}
	}
	return v1
}
