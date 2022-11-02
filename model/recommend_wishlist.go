package model

import "github.com/google/uuid"

type RecommendWishlist struct {
	RecommendId uuid.UUID `gorm:"type:uuid;primary_key;"`
	ResultName  string
	NeedToSave  int
	Status      string
	BalanceId   Balance `gorm:"foreignKey:WishlistId"`
}
