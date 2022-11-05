package model

import (
	"time"

	"gorm.io/gorm"
)

type Wishlist struct {
	WishlistId   string `gorm:"type:varchar(255);primary_key"`
	WhislistName string `gorm:"type:varchar(255)"`
	TargetMoney  uint   `gorm:"type:int(11)"`
	TargetMonth  uint   `gorm:"type:int(11)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	IsFinish     string         `gorm:"type:varchar(255)"`
	UserId       string         `gorm:"type:varchar(255);not null"`
	BalanceId    *Balance       `gorm:"foreignKey:WishlistId"`
}
