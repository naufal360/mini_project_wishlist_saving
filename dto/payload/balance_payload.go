package payload

type Balance struct {
	AmmountMoney int `gorm:"type:int(11)"`
}

type SavingMoney struct {
	SavingMoney int `json:"savingmoney" gorm:"type:int(11)" validate:"required"`
}
