package entity

import "time"

//book
type Order struct {
	OrderId      uint64    `gorm:"primary_key:auto_increment" json:"order_id"`
	CustomerName string    `gorm:"type:varchar(255)" json:"name"`
	OrderedAt    time.Time `gorm:"ordered_at" json:"ordered_at"`
	Item         []Item    `gorm:"foreignKey:OrderId;constraint:onUpdate:CASCADE,OnDelete:CASCADE;" json:"item"`
}
