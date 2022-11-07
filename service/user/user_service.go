package service

import (
	"errors"
	"wishlist/dto/payload"
	"wishlist/dto/response"
	m "wishlist/middleware"
	"wishlist/model"
	mysqldb "wishlist/repository/mysqldb/user"
	"wishlist/util"

	"github.com/google/uuid"
)

type UserService interface {
	RegisterUser(payload payload.Register) error
	LoginUser(payload payload.Login) (response.Login, error)
}

type userService struct {
	UserRepo mysqldb.UserRepository
}

func NewUserService(userRepo mysqldb.UserRepository) *userService {
	return &userService{
		UserRepo: userRepo,
	}
}

func (u *userService) RegisterUser(payload payload.Register) error {

	hashPass, err := util.HashPassword(payload.Password)
	if err != nil {
		return err
	}

	id := uuid.NewString()
	userModel := model.User{
		ID:       id,
		Name:     payload.Name,
		Username: payload.Username,
		Email:    payload.Email,
		Password: hashPass,
	}

	errRepo := u.UserRepo.RegisterUser(userModel)
	if errRepo != nil {
		return errors.New("username or email already exist")
	}

	return nil
}

func (u *userService) LoginUser(payload payload.Login) (response.Login, error) {
	var loginResponse response.Login

	userModel := model.User{
		Username: payload.Username,
		Password: payload.Password,
	}

	userData, err := u.UserRepo.LoginUser(userModel)
	if err != nil {
		return loginResponse, err
	}

	isValid := util.CheckPasswordHash(payload.Password, userData.Password)
	if !isValid {
		return loginResponse, errors.New("wrong password")
	}

	token, errToken := m.CreateToken(userData.ID, userData.Username)

	if errToken != nil {
		return loginResponse, err
	}

	loginResponse = response.Login{
		Token: token,
	}

	return loginResponse, nil
}
