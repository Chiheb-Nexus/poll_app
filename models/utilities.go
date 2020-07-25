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
	conn, err := gorm.Open("sqlite3", "../db.sqlite3")
	if err != nil {
		fmt.Printf("%v\n", err)
		panic("Cannot connect to DB!")
	}
	// FIXME: It will close at this point
	// You need to defer it in the main func
	// defer db.Close()
	// Set Stdout as backend logger
	conn.SetLogger(log.New(os.Stdout, "\r\n", 0))
	conn.LogMode(true)
	// Migrate the APP
	conn.AutoMigrate(&User{})
	conn.AutoMigrate(&Question{})
	conn.AutoMigrate(&Answer{})
	DB = conn
	return DB
}
