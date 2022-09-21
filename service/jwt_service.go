package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(UserID string) string
	ValidateToken(Token string) (*jwt.Token, error)
}

type jwtClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "fizahra",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRETKEY")
	if secretKey != "" {
		secretKey = "fiisecret"
	}
	return secretKey
}

func (js *jwtService) GenerateToken(UserID string) string {
	claim := &jwtClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    js.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tkn, err := token.SignedString([]byte(js.secretKey))
	if err != nil {
		panic(err)
	}
	return tkn
}

func (js *jwtService) ValidateToken(Token string) (*jwt.Token, error) {
	return jwt.Parse(Token, func(a_ *jwt.Token) (interface{}, error) {
		if _, ok := a_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v:", a_.Header["alg"])
		}
		return []byte(js.secretKey), nil
	})
}
