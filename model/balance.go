package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Balance struct {
	BalanceId    uuid.UUID `gorm:"type:uuid;primary_key;"`
	AmmountMoney int
	MonthSave    int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	WishlistId   Wishlist       `gorm:"foreignKey:WishlistId"`
}
