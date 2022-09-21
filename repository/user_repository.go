package repository

import (
	"assignment2go/entity"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	DuplicateEmail(email string) (tx *gorm.DB)
	VerifyCredential(email string, password string) interface{}
	FindUser(UserID string) entity.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(dbConn *gorm.DB) UserRepository {
	return &userConnection{
		connection: dbConn,
	}
}

func (uc *userConnection) InsertUser(user entity.User) entity.User {
	user.Password = HashandSalt([]byte(user.Password))
	uc.connection.Save(&user)
	return user
}

func (uc *userConnection) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = HashandSalt([]byte(user.Password))
	} else {
		var tempUser entity.User
		uc.connection.Find(&tempUser, user.UserID)
		user.Password = tempUser.Password
	}

	uc.connection.Save(&user)
	return user
}

func (uc *userConnection) DuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return uc.connection.Where("email = ?", email).Take(&user)

}

func (uc *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := uc.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (uc *userConnection) FindUser(UserID string) entity.User {
	var user entity.User
	uc.connection.Find(&user, UserID)
	return user
}

func HashandSalt(pass []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		log.Println("Failed to hash a password")
	}
	return string(hash)
}
