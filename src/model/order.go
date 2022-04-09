package model

import "time"

const (
	OrderStatusInCart    = 1
	OrderStatusPaid      = 2
	OrderStatusCancelled = 3
)

type Order struct {
	ID             int       `json:"id" gorm:"primaryKey"`
	Invoice        string    `json:"invoice"`
	CustomerName   string    `json:"customer_name"`
	BillAmount     float32   `json:"bill_amount"`
	TotalAmount    float32   `json:"total_amount"`
	DiscountAmount float32   `json:"discount_amount"`
	OrderStatus    int8      `json:"order_status" gorm:"size:1"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (Order) TableName() string {
	return "order"
}
