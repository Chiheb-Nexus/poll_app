package models

import (
	"fmt"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestUserCreation(t *testing.T) {

	db := Connector()
	defer db.Close()

	// Create User
	db.Create(&User{Username: "nexus", Age: 31})

	var user User
	// Find first User
	db.Where("username = ?", "nexus").First(&user)
	// Test created user
	if user.Username != "nexus" {
		msg := fmt.Sprintf("%s != nexus", user.Username)
		t.Fatal(msg)
	} else if user.Age != 31 {
		msg := fmt.Sprintf("%d != 31", user.Age)
		t.Fatal(msg)
	}
	// Update user
	db.Model(&user).Update("Username", "Nexus")
	db.Model(&user).Update("Age", 32)
	// Test updated user
	if user.Username != "Nexus" {
		msg := fmt.Sprintf("%s != nexus", user.Username)
		t.Fatal(msg)
	} else if user.Age != 32 {
		msg := fmt.Sprintf("%d != 31", user.Age)
		t.Fatal(msg)
	}

	var (
		question             Question
		userQuestionsCounter uint
	)

	// Create Question
	db.Create(&Question{TextQuestion: "toto", UserRefer: user.ID})
	db.First(&question)
	db.Model(&Question{}).Where("user_refer = ?", user.ID).Count(&userQuestionsCounter)
	if userQuestionsCounter != 1 {
		msg := fmt.Sprintf(
			"Counter: %d != 1 [UserID: %d | Question ID: %d]",
			userQuestionsCounter,
			user.ID,
			question.ID,
		)
		t.Fatal(msg)
	}

	// permanant delete
	db.Unscoped().Delete(&user)
	db.Unscoped().Delete(&question)

}
