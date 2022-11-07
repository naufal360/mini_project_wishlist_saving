package service

import (
	"errors"
	"math"
	"strconv"
	"wishlist/dto/payload"
	"wishlist/dto/response"
	m "wishlist/middleware"
	"wishlist/model"
	mysqldb "wishlist/repository/mysqldb/wishlist"

	"github.com/google/uuid"
)

type WishlistService interface {
	CreateWishlist(payload payload.Wishlist, auth string) error
	UpdateWishlist(payload payload.WishlistUpdate, id string) error
	UpdateBalance(payload payload.SavingMoney, auth, id string) error
	ReadWishlist(auth string) ([]model.Wishlist, error)
	ReadWishlistById(auth, id string) (model.Wishlist, error)
	ReadRecommend(auth, wishlistId string) (response.RecommendWishlist, error)
	DeleteWishlist(wishlistId string) error
}

type wishlistService struct {
	wishlistRepo mysqldb.WishlistRepository
}

func NewWishlistService(wishlistRepo mysqldb.WishlistRepository) *wishlistService {
	return &wishlistService{
		wishlistRepo: wishlistRepo,
	}
}

func (w *wishlistService) CreateWishlist(payload payload.Wishlist, auth string) error {

	userId, err := m.GetUserId(auth)
	if err != nil {
		return err
	}

	id := uuid.NewString()
	idBalance := uuid.NewString()
	idHistoryBalance := uuid.NewString()
	isFinish := "onprogress"
	statusHistory := "success"

	newData := model.Wishlist{
		WishlistId:   id,
		WhislistName: payload.WhislistName,
		TargetMoney:  payload.TargetMoney,
		TargetMonth:  payload.TargetMonth,
		IsFinish:     isFinish,
		UserId:       userId,
	}

	NewBalance := model.Balance{
		BalanceId:    idBalance,
		AmmountMoney: 0,
		ExceedMoney:  0,
		CountSave:    0,
		WishlistId:   id,
	}

	NewHistoryBalance := model.HistoryBalance{
		HistoryBalanceId: idHistoryBalance,
		SavingMoney:      0,
		Status:           statusHistory,
		BalanceIdHistory: idBalance,
	}

	err = w.wishlistRepo.CreateWishlist(newData)

	if err != nil {
		return err
	}

	err = w.wishlistRepo.CreateBalance(NewBalance)

	if err != nil {
		return err
	}

	err = w.wishlistRepo.CreateHistoryBalance(NewHistoryBalance)

	if err != nil {
		return err
	}
	return nil
}

func (w *wishlistService) UpdateWishlist(payload payload.WishlistUpdate, id string) error {
	newData := model.Wishlist{
		WishlistId:   id,
		WhislistName: payload.WhislistName,
		TargetMoney:  payload.TargetMoney,
		TargetMonth:  payload.TargetMonth,
	}

	if err := w.wishlistRepo.UpdateWishlist(newData); err != nil {
		return err
	}
	return nil
}

func (w *wishlistService) UpdateBalance(payload payload.SavingMoney, auth, id string) error {
	idHistoryBalance := uuid.NewString()

	userId, err := m.GetUserId(auth)
	if err != nil {
		return err
	}

	wishlist, err := w.wishlistRepo.ReadWishlistById(userId, id)

	if err != nil {
		return err
	}

	newMoney := wishlist.BalanceId.AmmountMoney + payload.SavingMoney
	exceedNewMoney := newMoney - int(wishlist.TargetMoney)
	countIncrement := wishlist.BalanceId.CountSave + 1
	statusHistory := "success"

	newData := model.Balance{
		BalanceId:    wishlist.BalanceId.BalanceId,
		AmmountMoney: newMoney,
		ExceedMoney:  0,
		CountSave:    countIncrement,
		WishlistId:   id,
	}

	if newMoney > int(wishlist.TargetMoney) {
		newData = model.Balance{
			BalanceId:    wishlist.BalanceId.BalanceId,
			AmmountMoney: int(wishlist.TargetMoney),
			ExceedMoney:  exceedNewMoney,
			CountSave:    countIncrement,
			WishlistId:   id,
		}
		if wishlist.IsFinish == "finish" {
			statusHistory = "failed"
		}
	}

	err = w.wishlistRepo.UpdateBalance(newData)
	if err != nil {
		return err
	}

	if newMoney >= int(wishlist.TargetMoney) {
		wishlistData := model.Wishlist{
			WishlistId: id,
			IsFinish:   "finish",
		}

		err := w.wishlistRepo.UpdateWishlist(wishlistData)
		if err != nil {
			return err
		}

		if err := w.wishlistRepo.DeleteWishlist(id); err != nil {
			return errors.New("youre wishlist already finish")
		}
	}

	NewHistoryBalance := model.HistoryBalance{
		HistoryBalanceId: idHistoryBalance,
		SavingMoney:      payload.SavingMoney,
		Status:           statusHistory,
		BalanceIdHistory: wishlist.BalanceId.BalanceId,
	}

	err = w.wishlistRepo.CreateHistoryBalance(NewHistoryBalance)

	if err != nil {
		return err
	}

	return nil
}

func (w *wishlistService) ReadWishlist(auth string) ([]model.Wishlist, error) {
	userId, err := m.GetUserId(auth)
	if err != nil {
		return []model.Wishlist{}, err
	}

	allWishlist, err := w.wishlistRepo.ReadWishlist(userId)

	if err != nil {
		return allWishlist, err
	}

	return allWishlist, nil
}

func (w *wishlistService) ReadWishlistById(auth, id string) (model.Wishlist, error) {
	userId, err := m.GetUserId(auth)
	if err != nil {
		return model.Wishlist{}, err
	}

	wishlist, err := w.wishlistRepo.ReadWishlistById(userId, id)

	if err != nil {
		return wishlist, err
	}

	return wishlist, nil
}

func (w *wishlistService) ReadRecommend(auth, wishlistId string) (response.RecommendWishlist, error) {

	var countRecommend, insufficient int

	userId, err := m.GetUserId(auth)
	if err != nil {
		return response.RecommendWishlist{}, err
	}

	wishlist, err := w.wishlistRepo.ReadWishlistById(userId, wishlistId)
	if err != nil {
		return response.RecommendWishlist{}, err
	}

	idealSaving := wishlist.TargetMoney / wishlist.TargetMonth
	insufficient = int(wishlist.TargetMoney) - wishlist.BalanceId.AmmountMoney

	if int(wishlist.TargetMonth) <= wishlist.BalanceId.CountSave && wishlist.BalanceId.AmmountMoney < int(wishlist.TargetMoney) {
		var nowIdeal float64 = float64(insufficient) / float64(idealSaving)
		countRecommend = int(math.Round(nowIdeal))
	} else {
		countRecommend = int(wishlist.TargetMonth) - wishlist.BalanceId.CountSave
	}

	var messageResponse string

	if insufficient <= 0 && countRecommend <= 0 {
		countRecommend = 0
		insufficient = 0
		messageResponse = "Congratulation, now you can buy your " + wishlist.WhislistName + "!"
	} else {
		messageResponse = "wishlist dengan nama " + wishlist.WhislistName + " serta dengan wishlist id " + wishlist.WishlistId +
			", mari menabung lagi sebanyak " + strconv.Itoa(countRecommend) +
			" kali dengan nominal sebesar Rp." + strconv.Itoa(insufficient/countRecommend) + ",- pada setiap kali menabung."
	}

	recommend := response.RecommendWishlist{
		WishlistId:        wishlist.WishlistId,
		Name:              wishlist.WhislistName,
		Insufficient:      insufficient,
		CountRecommend:    countRecommend,
		ResponseRecommend: messageResponse,
	}

	return recommend, nil
}

func (w *wishlistService) DeleteWishlist(wishlistId string) error {

	updateIsFinish := model.Wishlist{
		WishlistId: wishlistId,
		IsFinish:   "finish",
	}

	if err := w.wishlistRepo.UpdateWishlist(updateIsFinish); err != nil {
		return err
	}

	if err := w.wishlistRepo.DeleteWishlist(wishlistId); err != nil {
		return err
	}
	return nil
}
