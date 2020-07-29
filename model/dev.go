package model

import "github.com/jinzhu/gorm"

type Dev struct {
	gorm.Model
	Devid    string
	Password string
}
