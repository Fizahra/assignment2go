package dto

// type UserCreateDTO struct {
// 	Name     string `json:"name" form:"name" binding:"required"`
// 	Email    string `json:"email" form:"email" binding:"required,email"`
// 	Password string `json:"password,omiempty" form:"password,omiempty" binding:"required,min:6"`
// }

type UserUpdateDTO struct {
	UserID   uint64 `json:"user_id" form:"user_id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omiempty" form:"password,omiempty"`
}
