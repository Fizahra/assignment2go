package service

import (
	"assignment2go/dto"
	"assignment2go/entity"
	"assignment2go/repository"
	"log"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) entity.User
	DuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(UserRepo repository.UserRepository) AuthService {
	return &authService{
		userRepository: UserRepo,
	}
}

func (as *authService) VerifyCredential(email string, password string) interface{} {
	res := as.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparedPass := comparePass(v.Password, []byte(password))
		if v.Email == email && comparedPass {
			return res
		}
		return false
	}
	return false
}

func (as *authService) CreateUser(user dto.RegisterDTO) entity.User {
	userCreate := entity.User{}
	err := smapping.FillStruct(&userCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed to map %v", err)
	}
	res := as.userRepository.InsertUser(userCreate)
	return res
}

func (as *authService) DuplicateEmail(email string) bool {
	res := as.userRepository.DuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePass(hashedPass string, plainPass []byte) bool {
	byteHash := []byte(hashedPass)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPass)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
