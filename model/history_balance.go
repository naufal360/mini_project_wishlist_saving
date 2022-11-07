package model

import (
	"time"
)

type HistoryBalance struct {
	HistoryBalanceId string `gorm:"type:varchar(255);primary_key"`
	SavingMoney      int    `gorm:"type:int(11)"`
	Status           string `gorm:"type:varchar(255)"`
	CreatedAt        time.Time
	BalanceIdHistory string `gorm:"type:varchar(255)"`
}
