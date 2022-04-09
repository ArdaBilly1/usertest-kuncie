package repository

import (
	"gorm.io/gorm"
	"time"
	"usertest-kuncie/src/model"
)

type inventoryRepository struct {
	db *gorm.DB
}

type InventoryRepository interface {
	GetById(id int) (*model.Inventory, error)
	Insert(inventory *model.Inventory) error
	UpdateStock(id int, stock int32) error
}

func NewInventory(db *gorm.DB) InventoryRepository {
	return &inventoryRepository{db: db}
}

func (r *inventoryRepository) GetById(id int) (*model.Inventory, error) {
	inventory := new(model.Inventory)
	err := r.db.Model(inventory).
		Where("id", id).
		Find(&inventory).Error

	if err != nil {
		return &model.Inventory{}, err
	}

	return inventory, nil
}

func (r *inventoryRepository) Insert(inventory *model.Inventory) error {
	var data model.Inventory
	data.ItemID = inventory.ItemID
	data.Stock = inventory.Stock
	data.CreatedAt = time.Now()

	return r.db.Create(&data).Error
}

func (r *inventoryRepository) UpdateStock(id int, stock int32) error {
	var inventory model.Inventory
	err := r.db.Model(inventory).
		Where("id", id).
		Update("stock", stock).
		Error

	return err
}
