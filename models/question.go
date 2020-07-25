package models

import (
	"github.com/jinzhu/gorm"
)

// Question DB table
type Question struct {
	gorm.Model
	TextQuestion string
	UserRefer    uint
}
