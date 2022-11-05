package util

import (
	"wishlist/dto/payload"

	"github.com/go-playground/validator/v10"
)

func ValidateRegister(payload payload.Register) error {
	validate := validator.New()
	return validate.Struct(payload)
}

func ValidatePayloadWishlist(payload payload.Wishlist) error {
	validate := validator.New()
	return validate.Struct(payload)
}

func ValidateUpdatePayloadWishlist(payload payload.WishlistUpdate) error {
	validate := validator.New()
	return validate.Struct(payload)
}

func ValidateUpdatePayloadBalance(payload payload.SavingMoney) error {
	validate := validator.New()
	return validate.Struct(payload)
}
