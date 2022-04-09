package repository

import (
	"time"
	"usertest-kuncie/src/model"

	"gorm.io/gorm"
)

type OrderDetailRepository interface {
	Insert(orders []model.OrderDetail) error
}

type orderDetailRepository struct {
	db *gorm.DB
}

func NewOrderDetail(db *gorm.DB) OrderDetailRepository {
	return &orderDetailRepository{db: db}
}

func (r *orderDetailRepository) Insert(orders []model.OrderDetail) error {
	var details []model.OrderDetail

	for i := range orders {
		var detail model.OrderDetail
		detail.ItemId = orders[i].ItemId
		detail.OrderId = orders[i].OrderId
		detail.CapitalPrice = orders[i].CapitalPrice
		detail.Qty = orders[i].Qty
		detail.CreatedAt = time.Now()
		details = append(details, detail)
	}

	return r.db.Model(details).Create(&details).Error
}
