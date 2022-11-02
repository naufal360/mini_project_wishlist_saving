package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Name     string
	Username string
	Email    string
	Password string
}
