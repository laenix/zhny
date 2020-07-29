package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Devdata struct {
	gorm.Model
	Devid          string `gorm:"size:255;not null"`
	Devtemperature string `gorm:"size:20;not null"`
	Devhumidity    string `gorm:"size:20;not null"`
	Devco2         string `gorm:"size:20;not null"`
	Time           time.Time
}
