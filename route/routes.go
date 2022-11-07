package route

import (
	"wishlist/config"
	mid "wishlist/middleware"
	ur "wishlist/route/user"
	wr "wishlist/route/wishlist"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	dbConfig := config.InitGorm()

	userAPI := config.InitUserAPI(dbConfig)
	wishlistAPI := config.InitWishlistAPI(dbConfig)

	routes := echo.New()

	// trailing slash
	mid.RemoveSlash(routes)

	// set logger
	mid.LogMiddleware(routes)

	// Unauthenticated
	ur.UserUnauthenticated(routes, userAPI)

	// Authenticated

	// user
	ur.UserAuthenticated(routes, userAPI)
	// wishlist
	wr.WishlistAuthenticated(routes, wishlistAPI)

	return routes
}
