package models

import (
	"fmt"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestQuestionCreation(t *testing.T) {
	db := Connector()
	defer db.Close()

	// Create a user
	db.Create(&User{Username: "nexus", Age: 31})
	// Filter first User
	var user User
	db.First(&user)
	// Create Question
	db.Create(&Question{TextQuestion: "toto", UserRefer: user.ID})
	// Filter first Question
	var question Question
	db.First(&question)
	if question.TextQuestion != "toto" {
		msg := fmt.Sprintf("%s != toto", question.TextQuestion)
		t.Fatal(msg)
	}
	// permanant delete
	db.Unscoped().Delete(&question)
	db.Unscoped().Delete(&user)
}
