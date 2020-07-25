package models

import (
	"github.com/jinzhu/gorm"
)

// User DB table
type User struct {
	gorm.Model
	Username  string `gorm:"size=255"`
	Age       int
	Questions []Question `gorm:"foreignkey=UserID"`
}
