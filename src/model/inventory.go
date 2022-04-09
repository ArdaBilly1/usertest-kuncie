package model

import (
	"time"
)

type Inventory struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	ItemID    int       `json:"item_id"`
	Stock     int32     `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Inventory) TableName() string {
	return "inventory"
}
