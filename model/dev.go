package model

import "github.com/jinzhu/gorm"

type Devs struct {
	gorm.Model
	Devid   string `gorm:"size:255;not null"`
	Devpass string `gorm:"size:255;not null"`
	Belong  string `gorm:"size:255"`
	Cmd     string `gorm:"size:255"`
}
