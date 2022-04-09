package response

type ItemResponse struct {
	ID        int     `json:"id"`
	SKU       string  `json:"sku" validate:"required"`
	ItemName  string  `json:"item_name" validate:"required"`
	ItemPrice float32 `json:"item_price" validate:"required"`
}
