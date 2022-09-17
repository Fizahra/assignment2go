package entity

import "time"

//book
type Order struct {
	Order_Id      uint64    `gorm:"primary_key:auto_increment" json:"order_id"`
	Customer_Name string    `gorm:"type:varchar(255)" json:"name"`
	Ordered_At    time.Time `gorm:"ordered_at" json:"ordered_at"`
	Item          []Item    `gorm:"foreignKey:Order_Id;constraint:onUpdate:CASCADE,OnDelete:CASCADE;" json:"item"`
}
