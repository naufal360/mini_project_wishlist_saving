package user_route

import (
	"wishlist/controller"

	"github.com/labstack/echo/v4"
)

func UserUnauthenticated(routes *echo.Echo, api *controller.UserController) {
	user := routes.Group("/v1")
	{
		user.POST("/signup", api.RegisterUser)
		user.POST("/login", api.LoginUser)

	}
}
