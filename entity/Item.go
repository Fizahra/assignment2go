package entity

//user
type Item struct {
	Item_Id     uint64 `gorm:"primary_key:auto_increment" json:"item_id"`
	Item_Code   uint64 `gorm:"type:varchar(255)" json:"item_code"`
	Order_Id    uint64 `gorm:"column:order_id" json:"order_id"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Quantity    uint64 `gorm:"type:int(11)" json:"quantity"`
}
