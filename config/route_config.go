package config

import (
	"wishlist/controller"
	mysqldb_u "wishlist/repository/mysqldb/user"
	mysqldb_w "wishlist/repository/mysqldb/wishlist"
	service "wishlist/service/user"
	service_w "wishlist/service/wishlist"

	"gorm.io/gorm"
)

func InitUserAPI(db *gorm.DB) *controller.UserController {
	userRepo := mysqldb_u.NewUserRepository(db)
	userServ := service.NewUserService(userRepo)
	userAPI := controller.NewUserController(userServ)
	return userAPI
}

func InitWishlistAPI(db *gorm.DB) *controller.WishlistController {
	wishlistRepo := mysqldb_w.NewWishlistRepository(db)
	wishlistServ := service_w.NewWishlistService(wishlistRepo)
	wishlistAPI := controller.NewWishlistController(wishlistServ)
	return wishlistAPI
}
