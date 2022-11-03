package postgresdb

import (
	"wishlist/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(data model.User) error
	LoginUser(data model.User) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (u *userRepository) RegisterUser(data model.User) error {
	if err := u.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) LoginUser(data model.User) (model.User, error) {
	var user model.User

	createData := u.db.Where("username = ?", data.Username).First(&user)
	if err := createData.Error; err != nil {
		return user, err
	}
	return user, nil
}
