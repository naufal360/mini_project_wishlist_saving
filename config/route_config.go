package config

import (
	"wishlist/controller"
	mysqldb "wishlist/repository/mysqldb/user"
	service "wishlist/service/user"

	"gorm.io/gorm"
)

func InitUserAPI(db *gorm.DB) *controller.UserController {
	userRepo := mysqldb.NewUserRepository(db)
	userServ := service.NewUserService(userRepo)
	userAPI := controller.NewUserController(userServ)
	return userAPI
}
