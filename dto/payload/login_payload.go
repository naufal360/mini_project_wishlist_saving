package payload

type Login struct {
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}
