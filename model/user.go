package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"type:varchar(255);primary_key"`
	Name      string `gorm:"type:varchar(255)"`
	Username  string `gorm:"type:varchar(255);unique"`
	Email     string `gorm:"type:varchar(100);unique"`
	Password  string `gorm:"type:varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
