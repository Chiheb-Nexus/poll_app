package main

import (
	"github.com/chiheb-nexus/poll_app/controllers"
	"github.com/chiheb-nexus/poll_app/models"
	"github.com/chiheb-nexus/poll_app/utilities"
	"github.com/gin-gonic/gin"

	// Need to add it for gorm DB connection
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var versionAPI = "v1"

func main() {
	r := gin.Default()
	DB := models.Connector()
	defer DB.Close()
	r.GET(utilities.FormatAPIVersion("users", versionAPI), controllers.AllUsers)
	r.GET(utilities.FormatAPIVersion("questions", versionAPI), controllers.AllQuestions)
	r.Run()
}
