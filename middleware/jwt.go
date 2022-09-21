package middleware

import (
	"assignment2go/helper"
	"assignment2go/service"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		autHeader := c.GetHeader("Authorization")
		if autHeader == "" {
			res := helper.BuildErrorResponse("Failed to process your request", "There's no token found!", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
		}
		token, err := jwtService.ValidateToken(autHeader)
		if token.Valid {
			claim := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claim["user_id"])
			log.Println("Claim[issuer]: ", claim["issuer"])
		} else {
			log.Println(err)
			res := helper.BuildErrorResponse("Token isn't valid!", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		}
	}
}
