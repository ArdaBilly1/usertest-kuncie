package request

type PromoInsertRequest struct {
	Name        string  `json:"name" validate:"required"`
	Type        int8    `json:"type" validate:"required"`
	MinQty      int16   `json:"min_qty" validate:"required"`
	ItemId      int     `json:"item_id"`
	ItemBonusId int     `json:"item_bonus_id"`
	Percentage  float32 `json:"percentage"`
}
