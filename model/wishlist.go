package model

import (
	"time"

	"gorm.io/gorm"
)

type Wishlist struct {
	WishlistId   string `gorm:"type:varchar(255);primary_key"`
	WhislistName string
	TargetMoney  string
	TargetMonth  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	UserId       User           `gorm:"foreignKey:ID"`
}
