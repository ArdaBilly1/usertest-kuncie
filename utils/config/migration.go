package config

import (
	"gorm.io/gorm"
	"usertest-kuncie/src/model"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(
		model.Inventory{},
		model.Item{},
		model.Order{},
		model.OrderDetail{},
		model.Promo{},
	)
}
