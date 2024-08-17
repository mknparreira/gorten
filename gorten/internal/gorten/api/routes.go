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
			userRoute := handlers.User()
			users.GET("/", userRoute.List)
			users.GET("/:id", userRoute.UserByID)
			users.PUT("/:id", userRoute.UpdateByID)
			users.POST("/", userRoute.Create)
		}
	}
}
