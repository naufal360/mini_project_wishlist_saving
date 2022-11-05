package model

import (
	"time"
)

type Balance struct {
	BalanceId       string `gorm:"type:varchar(255);primary_key"`
	AmmountMoney    int    `gorm:"type:int(11)"`
	CountSave       int    `gorm:"type:int(11)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	WishlistId      string           `gorm:"type:varchar(255);not null"`
	HistoryBalances []HistoryBalance `gorm:"foreignKey:BalanceIdHistory"`
}
