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
		authUser.GET("/wishlists", api.ReadWishlist)
		authUser.GET("/wishlists/:wishlistid", api.ReadWishlistById)
		authUser.GET("/wishlists/:wishlistid/recommend", api.ReadRecommend)
		authUser.POST("/wishlists", api.CreateWishlist)
		authUser.PUT("/wishlists/:wishlistid", api.UpdateWishlist)
		authUser.PUT("/wishlists/:wishlistid/balances", api.UpdateBalance)
		authUser.DELETE("/wishlists/:wishlistid", api.DeleteWishlist)
	}
}
