package entity

//struct order yang bakal dijadiin tabel database
type User struct {
	UserID   uint64 `gorm:"primary_key:auto_increment" json:"user_id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"uniqueIndex; type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null*" json:"-"`
	Token    string `gorm:"-" json:"token,omitempty"`
}
