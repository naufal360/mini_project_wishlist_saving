package model

import (
	"time"
)

type HistoryBalance struct {
	HistoryBalanceId string `gorm:"type:varchar(255);primary_key"`
	Status           string
	CreatedAt        time.Time
	BalanceId        Balance `gorm:"foreignKey:BalanceId"`
}
