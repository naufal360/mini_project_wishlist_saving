package user_route

import (
	"net/http"
	"os"
	"wishlist/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserUnauthenticated(routes *echo.Echo, api *controller.UserController) {
	user := routes.Group("/v1")
	{
		user.POST("/signup", api.RegisterUser)
		user.POST("/login", api.LoginUser)

	}
}

func UserAuthenticated(routes *echo.Echo, api *controller.UserController) {
	authUser := routes.Group("/v1")
	authUser.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT_KEY"))))

	{
		authUser.GET("/user", func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]string{
				"test": "testuser",
			})
		})
	}
}
