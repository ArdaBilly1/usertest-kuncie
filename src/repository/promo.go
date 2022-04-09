package repository

import (
	"gorm.io/gorm"
	"time"
	"usertest-kuncie/src/model"
)

type PromoRepository interface {
	GetById(id int) (promo model.Promo, err error)
	Insert(promo model.Promo) error
}
type promoRepository struct {
	db *gorm.DB
}

func NewPromo(db *gorm.DB) PromoRepository {
	return &promoRepository{db: db}
}

func (r *promoRepository) GetById(id int) (promo model.Promo, err error) {
	err = r.db.Model(promo).
		Where("id", id).
		Find(&promo).
		Error
	return
}

func (r *promoRepository) Insert(promo model.Promo) error {
	var data model.Promo

	data.ID = promo.ID
	data.ItemBonusId = promo.ItemBonusId
	data.Percentage = promo.Percentage
	data.Name = promo.Name
	data.MinQty = promo.MinQty
	data.Type = promo.Type
	data.CreatedAt = time.Now()

	return r.db.Model(data).Create(&data).Error
}
