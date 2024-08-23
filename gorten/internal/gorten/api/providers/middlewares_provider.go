package providers

import (
	"gorten/internal/gorten/api/middlewares"

	"github.com/gin-gonic/gin"
)

func MiddlewaresProvider() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.ErrorHandlerMiddleware())
	return r
}
