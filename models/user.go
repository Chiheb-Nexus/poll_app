package models

import "github.com/jinzhu/gorm"

// User DB table
type User struct {
	gorm.Model
	Username  string     `json:"username" gorm:"size=255"`
	Age       int        `json:"age"`
	Questions []Question `json:"questions" gorm:"foreignkey=UserID"`
}
