package repository

import (
	"time"
	"usertest-kuncie/src/model"

	"gorm.io/gorm"
)

type ItemRepository interface {
	GetAll() (items []model.Item, err error)
	GetById(id int) (item model.Item, err error)
	Create(item model.Item) error
	Update(item model.Item, id int) error
	GetSKU(sku string) (id int, err error)
}

type itemRepository struct {
	db *gorm.DB
}

func NewItem(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) GetAll() (items []model.Item, err error) {
	var item model.Item

	err = r.db.Model(item).
		Find(&items).Error

	return
}

func (r *itemRepository) GetById(id int) (item model.Item, err error) {
	err = r.db.Model(item).
		Where("id", id).
		Find(&item).
		Error

	return
}

func (r *itemRepository) Create(item model.Item) error {
	var data model.Item

	data.ItemName = item.ItemName
	data.SKU = item.SKU
	data.ItemPrice = item.ItemPrice
	data.CreatedAt = time.Now()

	err := r.db.Create(&data).Error
	return err
}

func (r *itemRepository) Update(item model.Item, id int) error {
	return r.db.Model(item).
		Where("id", id).
		Updates(&item).
		Error
}

func (r *itemRepository) GetSKU(sku string) (id int, err error) {
	var item model.Item
	instance := r.db.Model(item).
		Where("sku", sku).
		Limit(1).
		Order("item.created_at ASC").
		Row()
	err = instance.Scan(id)
	return
}
