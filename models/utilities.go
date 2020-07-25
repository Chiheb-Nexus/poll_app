package models

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

// Connector DB
func Connector() *gorm.DB {
	db, err := gorm.Open("sqlite3", "../sqlite3.db")
	if err != nil {
		panic("Cannot connect to DB!")
	}
	// FIXME: It will close at this point
	// You need to defer it in the main func
	// defer db.Close()
	// Set Stdout as backend logger
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	return db
}
