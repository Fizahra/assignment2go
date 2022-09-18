package entity

//user
type Item struct {
	ItemId      uint64 `gorm:"primary_key:auto_increment" json:"item_id"`
	ItemCode    uint64 `gorm:"type:varchar(255)" json:"item_code"`
	OrderId     uint64 `gorm:"column:order_id" json:"order_id"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Quantity    uint64 `gorm:"type:int(11)" json:"quantity"`
}
