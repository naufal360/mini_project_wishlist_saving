package model

import (
	"time"

	"gorm.io/gorm"
)

type Balance struct {
	BalanceId    string `gorm:"type:varchar(255);primary_key"`
	AmmountMoney int
	MonthSave    int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	WishlistId   Wishlist       `gorm:"foreignKey:WishlistId"`
}
