package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	email        string
	username     string
	passwordhash string
	name         string
	lastname     string
	role         string
}
