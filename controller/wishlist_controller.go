package controller

import (
	"net/http"
	"wishlist/dto/payload"
	service "wishlist/service/wishlist"
	"wishlist/util"

	"github.com/labstack/echo/v4"
)

type WishlistController struct {
	WishlistService service.WishlistService
}

func NewWishlistController(wishlistServ service.WishlistService) *WishlistController {
	return &WishlistController{
		WishlistService: wishlistServ,
	}
}

func (w *WishlistController) CreateWishlist(ctx echo.Context) error {
	var payload payload.Wishlist

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := util.ValidatePayloadWishlist(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	auth := ctx.Request().Header.Get("Authorization")

	err := w.WishlistService.CreateWishlist(payload, auth)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success create wishlist",
	})
}

func (w *WishlistController) UpdateWishlist(ctx echo.Context) error {
	var payload payload.WishlistUpdate

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	id := ctx.Param("wishlistid")

	if err := util.ValidateUpdatePayloadWishlist(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := w.WishlistService.UpdateWishlist(payload, id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success updated wishlist",
	})
}

func (w *WishlistController) ReadWishlist(ctx echo.Context) error {
	auth := ctx.Request().Header.Get("Authorization")
	dataWishlist, err := w.WishlistService.ReadWishlist(auth)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all wishlist",
		"data":     dataWishlist,
	})
}

func (w *WishlistController) ReadRecommend(ctx echo.Context) error {
	wishlistId := ctx.Param("wishlistid")
	dataRecommend, err := w.WishlistService.ReadRecommend(wishlistId)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get recommendation",
		"data":     dataRecommend,
	})
}

func (w *WishlistController) UpdateBalance(ctx echo.Context) error {
	var payload payload.SavingMoney

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	id := ctx.Param("wishlistid")

	if err := util.ValidateUpdatePayloadBalance(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := w.WishlistService.UpdateBalance(payload, id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"messages": "saving success",
	})
}

func (w *WishlistController) DeleteWishlist(ctx echo.Context) error {
	wishlistId := ctx.Param("wishlistid")

	if err := w.WishlistService.DeleteWishlist(wishlistId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Delete wishlist success",
	})
}
