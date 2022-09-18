package dto

type ItemUpdateDTO struct {
	ItemId      uint64 `json:"id" form:"id" binding:"required"`
	ItemCode    uint64 `json:"item_code" form:"item_code" binding:"required"`
	OrderId     uint64 `json:"order_id" form:"order_id" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Quantity    uint64 `json:"quantity" form:"quantity" binding:"required"`
}

type ItemCreateDTO struct {
	ItemCode    uint64 `json:"item_code" form:"item_code" binding:"required"`
	OrderId     uint64 `json:"order_id" form:"order_id" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Quantity    uint64 `json:"quantity" form:"quantity" binding:"required"`
}
