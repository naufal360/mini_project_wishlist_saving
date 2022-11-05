package payload

type Wishlist struct {
	WhislistName string `json:"wishlistname" gorm:"size:255;not null" validate:"required"`
	TargetMoney  uint   `json:"targetmoney" gorm:"size:11;not null" validate:"required"`
	TargetMonth  uint   `json:"targetmonth" gorm:"size:11;not null" validate:"required"`
}

type WishlistUpdate struct {
	WhislistName string `json:"wishlistname" gorm:"size:255" validate:"required"`
	TargetMoney  uint   `json:"targetmoney" gorm:"size:11" validate:"required"`
	TargetMonth  uint   `json:"targetmonth" gorm:"size:11" validate:"required"`
}
