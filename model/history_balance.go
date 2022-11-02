package model

import (
	"time"

	"github.com/google/uuid"
)

type HistoryBalance struct {
	HistoryBalanceId uuid.UUID `gorm:"type:uuid;primary_key;"`
	Status           string
	CreatedAt        time.Time
	BalanceId        Balance `gorm:"foreignKey:BalanceId"`
}
