package models

import (
	"fmt"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestAnswerCreation(t *testing.T) {
	db := Connector()
	defer db.Close()
	db.AutoMigrate(&Answer{})

	var (
		user     User
		question Question
		answer   Answer
	)
	// Create User & Question
	db.Create(&User{Username: "nexus", Age: 31})
	db.First(&user)
	db.Create(&Question{TextQuestion: "toto", UserRefer: user.ID})
	db.First(&question)

	// Create Answer
	db.Create(&Answer{TextAnswer: "hola", UserRef: user.ID, QuestionRef: question.ID})
	db.First(&answer)
	// Test Answer
	if answer.TextAnswer != "hola" {
		msg := fmt.Sprintf("%s != hola", answer.TextAnswer)
		t.Fatal(msg)
	} else if answer.UserRef != user.ID {
		msg := fmt.Sprintf("%d != %d", answer.UserRef, user.ID)
		t.Fatal(msg)
	} else if answer.QuestionRef != question.ID {
		msg := fmt.Sprintf("%d != %d", answer.QuestionRef, question.ID)
		t.Fatal(msg)
	}

	// Permanent delete
	db.Unscoped().Delete(&question)
	db.Unscoped().Delete(&user)
	db.Unscoped().Delete(&answer)

}
