package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GeneralHandlerImpl interface {
	PingHandler(c *gin.Context)
}

type GeneralHandler struct{}

func (h *GeneralHandler) PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
