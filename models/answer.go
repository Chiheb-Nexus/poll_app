package models

import "github.com/jinzhu/gorm"

// Answer DB table
type Answer struct {
	gorm.Model
	TextAnswer  string
	UserRef     uint
	QuestionRef uint
}
