package model

import "time"

const (
	PromoTypePercentage = 1
	PromoTypeBonusItem  = 2
)

type Promo struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Type        int8      `json:"type"`
	MinQty      int16     `json:"min_qty"`
	ItemBonusId int       `json:"item_bonus_id"`
	Percentage  float32   `json:"percentage"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Promo) TableName() string {
	return "promo"
}
