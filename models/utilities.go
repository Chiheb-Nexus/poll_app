package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

// DB connection
var DB *gorm.DB

// Connector DB
func Connector() *gorm.DB {
	// db.sqlite3 will be at the same level of main.go
	db, err := gorm.Open("sqlite3", "db.sqlite3")
	if err != nil {
		fmt.Printf("%v\n", err)
		panic("Cannot connect to DB!")
	}
	// FIXME: It will close at this point
	// You need to defer it in the main func
	// defer db.Close()
	// Set Stdout as backend logger
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	db.LogMode(true)
	// Migrate the APP
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Question{})
	db.AutoMigrate(&Answer{})
	DB = db
	return DB
}
