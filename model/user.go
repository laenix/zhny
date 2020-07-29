package model

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Password string `gorm:"size:255;not null"`
}
