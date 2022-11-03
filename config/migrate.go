package config

import (
	"wishlist/model"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
		// model.Wishlist{},
		// model.Balance{},
		// model.HistoryBalance{},
	)
}
