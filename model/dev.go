package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Devs struct {
	gorm.Model
	Devid   string `gorm:"size:255;not null"`
	Devpass string `gorm:"size:255;not null"`
	Belong  string `gorm:"size:255"`
	Cmd     string `gorm:"size:255"`
}

type Dev struct {
	Devid          string
	Belong         string
	Cmd            string
	Devtemperature string
	Devhumidity    string
	Devco2         string
	Time           time.Time
}
