package postgresdb

import (
	"wishlist/model"

	"gorm.io/gorm"
)

type WishlistRepository interface {
	CreateWishlist(data model.Wishlist) error
	CreateBalance(data model.Balance) error
	CreateHistoryBalance(data model.HistoryBalance) error
	UpdateWishlist(data model.Wishlist) error
	UpdateBalance(data model.Balance) error
	ReadWishlist(userId string) ([]model.Wishlist, error)
	ReadWishlistById(auth, id string) (model.Wishlist, error)
	ReadBalanceById(id string) (model.Balance, error)
	DeleteWishlist(wishlistId string) error
	DeleteBalanced(balancedId string) error
}

type wishlistRepository struct {
	db *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) *wishlistRepository {
	return &wishlistRepository{db: db}
}

func (w *wishlistRepository) CreateWishlist(data model.Wishlist) error {
	if err := w.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (w *wishlistRepository) CreateBalance(data model.Balance) error {
	if err := w.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (w *wishlistRepository) CreateHistoryBalance(data model.HistoryBalance) error {
	if err := w.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (w *wishlistRepository) UpdateWishlist(data model.Wishlist) error {
	if err := w.db.Model(&model.Wishlist{}).Where("wishlist_id = ?", data.WishlistId).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (w *wishlistRepository) UpdateBalance(data model.Balance) error {
	if err := w.db.Model(&model.Balance{}).Where("balance_id = ?", data.BalanceId).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (w *wishlistRepository) ReadWishlist(userId string) ([]model.Wishlist, error) {
	var wishlist []model.Wishlist
	if err := w.db.Unscoped().Preload("BalanceId").Preload("BalanceId.HistoryBalances").Where("user_id = ?", userId).Find(&wishlist).Error; err != nil {
		return wishlist, err
	}
	return wishlist, nil
}

func (w *wishlistRepository) ReadWishlistById(auth, id string) (model.Wishlist, error) {
	var wishlist model.Wishlist
	if err := w.db.Unscoped().Preload("BalanceId").Preload("BalanceId.HistoryBalances").Where("wishlist_id = ? AND user_id = ?", id, auth).Find(&wishlist).Error; err != nil {
		return wishlist, err
	}
	return wishlist, nil
}

func (w *wishlistRepository) ReadBalanceById(id string) (model.Balance, error) {
	var data model.Balance
	if err := w.db.Model(&model.Balance{}).Where("balance_id = ?", id).First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (w *wishlistRepository) DeleteWishlist(wishlistId string) error {
	var data model.Wishlist
	if err := w.db.Where("wishlist_id = ?", wishlistId).Find(&data).Delete(&data).Error; err != nil {
		return err
	}
	return nil
}

func (w *wishlistRepository) DeleteBalanced(balancedId string) error {
	var data model.Balance
	if err := w.db.Where("balance_id = ?", balancedId).First(&data).Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
