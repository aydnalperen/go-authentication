package models

import "gorm.io/gorm"

type User_no_gorm struct {
	email        string
	username     string
	passwordhash string
	firstname    string
	lastname     string
	role         string
}
type User struct {
	gorm.Model
	email        string
	username     string
	passwordhash string
	firstname    string
	lastname     string
	role         string
}

func (user *User) PairBody(u *User_no_gorm) {

	user.email = u.email
	user.username = u.username
	user.passwordhash = u.passwordhash
	user.firstname = u.firstname
	user.lastname = u.lastname
	user.role = u.role
}
