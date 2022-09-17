package dto

type ItemUpdateDTO struct {
	Item_Id     uint64 `json:"id" form:"id" binding:"required"`
	Item_Code   uint64 `json:"item_code" form:"item_code" binding:"required"`
	Order_Id    uint64 `json:"order_id" form:"order_id" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Quantity    uint64 `json:"quantity" form:"quantity" binding:"required"`
}

type ItemCreateDTO struct {
	Item_Code   uint64 `json:"item_code" form:"item_code" binding:"required"`
	Order_Id    uint64 `json:"order_id" form:"order_id" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Quantity    uint64 `json:"quantity" form:"quantity" binding:"required"`
}
