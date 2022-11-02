package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Wishlist struct {
	WishlistId   uuid.UUID `gorm:"type:uuid;primary_key;"`
	WhislistName string
	TargetMoney  string
	TargetMonth  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	UserId       User           `gorm:"foreignKey:Id"`
}
