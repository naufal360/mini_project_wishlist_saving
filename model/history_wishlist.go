package model

import (
	"time"

	"github.com/google/uuid"
)

type HistoryWishlist struct {
	HistoryWishlistId uuid.UUID `gorm:"type:uuid;primary_key;"`
	Status            string
	CreatedAt         time.Time
	BalanceId         Balance `gorm:"foreignKey:BalanceId"`
	UserId            User    `gorm:"foreignKey:Id"`
}
