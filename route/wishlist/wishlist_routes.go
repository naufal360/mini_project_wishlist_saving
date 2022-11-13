package routes

import (
	"os"
	"wishlist/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func WishlistAuthenticated(routes *echo.Echo, api *controller.WishlistController) {
	authUser := routes.Group("/v1")
	authUser.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT_KEY"))))

	{
		authUser.GET("/wishlist", api.ReadWishlist)
		authUser.GET("/wishlist/:wishlistid", api.ReadWishlistById)
		authUser.GET("/wishlist/:wishlistid/recommend", api.ReadRecommend)
		authUser.POST("/wishlist", api.CreateWishlist)
		authUser.PUT("/wishlist/:wishlistid", api.UpdateWishlist)
		authUser.PUT("/wishlist/:wishlistid/balance", api.UpdateBalance)
		authUser.DELETE("/wishlist/:wishlistid", api.DeleteWishlist)
	}
}
