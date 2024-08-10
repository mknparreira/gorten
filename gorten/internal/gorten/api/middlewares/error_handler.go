package middlewares

import (
	"gorten/internal/gorten/models"
	"gorten/pkg/errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)

				c.JSON(http.StatusInternalServerError, models.ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: errors.ErrInternalServerError.Error(),
				})
				c.Abort()
			}
		}()

		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			c.JSON(c.Writer.Status(), models.ErrorResponse{
				Code:    c.Writer.Status(),
				Message: err.Error(),
			})
		}
	}
}
