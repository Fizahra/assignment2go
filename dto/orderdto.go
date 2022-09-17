package dto

import "time"

type OrderUpdateDTO struct {
	Order_Id      uint64          `json:"id" form:"order_id" binding:"required"`
	Customer_Name string          `json:"name" form:"name" binding:"required"`
	Ordered_At    time.Time       `json:"ordered_at" form:"ordered_at" binding:"required"`
	Item          []ItemUpdateDTO `json:"item" form:"item" binding:"required"`
}

type OrderCreateDTO struct {
	Customer_Name string          `json:"name" form:"name" binding:"required"`
	Ordered_At    time.Time       `json:"ordered_at" form:"ordered_at" binding:"required"`
	Item          []ItemCreateDTO `json:"item" form:"item" binding:"required"`
}
