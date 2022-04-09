package model

import "time"

type OrderDetail struct {
	ID           int       `json:"id"`
	OrderId      int       `json:"order_id"`
	ItemId       int       `json:"item_id"`
	Price        float32   `json:"price"`
	CapitalPrice float32   `json:"capital_price"`
	Qty          int8      `json:"qty"`
	CreatedAt    time.Time `json:"created_at"`
}

func (OrderDetail) TableName() string {
	return "order_detail"
}
