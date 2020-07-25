package controllers

import (
	"net/http"

	"github.com/chiheb-nexus/poll_app/models"
	"github.com/gin-gonic/gin"
)

// AllUsers an API that returns all DB users with their informations
func AllUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}
