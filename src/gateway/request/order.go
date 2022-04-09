package request

type OrderRequest struct {
	CustomerName string               `json:"customer_name" validate:"required"`
	OrderDetail  []OrderDetailRequest `json:"order_detail" validate:"required"`
}

type OrderDetailRequest struct {
	ItemId int  `json:"item_id" validate:"required"`
	Qty    int8 `json:"qty" validate:"required"`
}
