package model

import (
	"time"

	"gorm.io/gorm"
)


type Recipe struct {
	gorm.Model
	Title       string
	Description string
	PrepTime    int `gorm:"column:prep_time"` 
	Servings    int
	PublishDate time.Time `gorm:"column:publish_date"`
	UserID      uint     
	User        User      `gorm:"foreignKey:UserID"` 
}
