package payload

type Register struct {
	Name     string `json:"name" gorm:"size:255;not null" validate:"required"`
	Username string `json:"username" gorm:"size:255;not null" validate:"required,min=4"`
	Email    string `json:"email" gorm:"size:100;not null" validate:"required,email"`
	Password string `json:"password" gorm:"size:100;not null" validate:"required,min=6"`
}
