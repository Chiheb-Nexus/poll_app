package main

import (
	"github.com/chiheb-nexus/poll_app/controllers"
	"github.com/chiheb-nexus/poll_app/models"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	r := gin.Default()
	_ = models.Connector()
	r.GET("/users", controllers.AllUsers)
	r.Run()
}
