package dto

import "time"

type OrderUpdateDTO struct {
	OrderId      uint64          `json:"id" form:"order_id" binding:"required"`
	CustomerName string          `json:"name" form:"name" binding:"required"`
	OrderedAt    time.Time       `json:"ordered_at" form:"ordered_at" binding:"required"`
	Item         []ItemUpdateDTO `json:"item" form:"item" binding:"required"`
	UserID       uint64          `json:"user_id" form:"user_id" binding:"required"`
}

type OrderCreateDTO struct {
	CustomerName string          `json:"name" form:"name" binding:"required"`
	OrderedAt    time.Time       `json:"ordered_at" form:"ordered_at" binding:"required"`
	Item         []ItemCreateDTO `json:"item" form:"item" binding:"required"`
	UserID       uint64          `json:"user_id" form:"user_id" binding:"required"`
}
