package repository

import (
	"time"
	"usertest-kuncie/src/model"

	"gorm.io/gorm"
)

type OrderRepositoryContract interface {
	Insert(order model.Order) error
}

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrder(DB *gorm.DB) OrderRepositoryContract {
	return &OrderRepository{DB: DB}
}

func (r *OrderRepository) Insert(order model.Order) error {
	var data model.Order

	data.Invoice = order.Invoice
	data.CustomerName = order.CustomerName
	data.BillAmount = order.BillAmount
	data.DiscountAmount = order.DiscountAmount
	data.OrderStatus = order.OrderStatus
	data.TotalAmount = order.TotalAmount
	data.CreatedAt = time.Now()

	return r.DB.Model(data).Create(&data).Error
}
