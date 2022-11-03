package route

import (
	"wishlist/config"
	mid "wishlist/middleware"
	ur "wishlist/route/user"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	dbConfig := config.InitGorm()

	userAPI := config.InitUserAPI(dbConfig)

	routes := echo.New()

	// trailing slash
	mid.RemoveSlash(routes)

	// set logger
	mid.LogMiddleware(routes)

	// Unauthenticated
	ur.UserUnauthenticated(routes, userAPI)

	// Authenticated
	ur.UserAuthenticated(routes, userAPI)

	return routes
}
