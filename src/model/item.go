package model

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	SKU       string         `json:"sku"`
	ItemName  string         `json:"item_name"`
	ItemPrice float32        `json:"item_price"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (Item) TableName() string {
	return "item"
}
