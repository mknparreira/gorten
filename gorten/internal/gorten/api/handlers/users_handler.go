package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func User() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "list of users"})
}

func (h *UserHandler) UserByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "listing a specific user"})
}

func (h *UserHandler) Create(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func (h *UserHandler) UpdateByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "user update"})
}
