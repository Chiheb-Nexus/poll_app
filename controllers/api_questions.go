package controllers

import (
	"net/http"

	"github.com/chiheb-nexus/poll_app/models"
	"github.com/gin-gonic/gin"
)

// AllQuestions API that returns all questions in DB
func AllQuestions(c *gin.Context) {
	var questions []models.Question
	models.DB.Find(&questions)
	c.JSON(http.StatusOK, gin.H{"data": questions})
}
