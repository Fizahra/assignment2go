package service

import (
	"assignment2go/dto"
	"assignment2go/entity"
	"assignment2go/repository"
	"log"

	"github.com/mashingan/smapping"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	FindUser(UserID string) entity.User
}

type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		UserRepository: ur,
	}
}

func (us *userService) Update(user dto.UserUpdateDTO) entity.User {
	userup := entity.User{}
	err := smapping.FillStruct(&userup, smapping.MapFields(user))
	if err != nil {
		log.Fatalf("Failed to find user! %v", err)
	}
	upuser := us.UserRepository.UpdateUser(userup)
	return upuser
}

func (us *userService) FindUser(UserID string) entity.User {
	return us.UserRepository.FindUser(UserID)
}
