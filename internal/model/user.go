package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username" binding:"required"`
	Email    string `gorm:"unique" json:"email" binding:"required,email"`
}
