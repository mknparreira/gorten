package middlewares

import (
	"errors"
	"gorten/internal/gorten/models"
	pkgerr "gorten/pkg/errors"
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
					Message: pkgerr.ErrInternalServerError.Error(),
				})
				c.Abort()
			}
		}()

		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			statusCode, errorMessage := mapErrors(err)

			c.JSON(statusCode, models.ErrorResponse{
				Code:    statusCode,
				Message: errorMessage,
			})
			c.Abort()
		}
	}
}

func mapErrors(err error) (int, string) {
	var customErr *pkgerr.CustomError
	if errors.As(err, &customErr) {
		return customErr.StatusCode, customErr.Message
	}

	return http.StatusInternalServerError, "An unexpected error occurred"
}
