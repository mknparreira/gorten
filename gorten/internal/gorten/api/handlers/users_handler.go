package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "list of users"})
}

func ListUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "listing a specific user"})
}

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "user updated"})
}
