package controller

import (
	"assignment2go/dto"
	"assignment2go/helper"
	"assignment2go/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Update(c *gin.Context)
	FindUser(c *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (uc *userController) Update(c *gin.Context) {
	var user dto.UserUpdateDTO
	err := c.ShouldBind(&user)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to update user", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	autHeader := c.GetHeader("Authorization")
	token, errt := uc.jwtService.ValidateToken(autHeader)
	if errt != nil {
		panic(errt.Error())
	}
	claim := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claim["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	user.UserID = id
	u := uc.userService.Update(user)
	res := helper.BuildResponse(true, "OK", u)
	c.JSON(http.StatusOK, res)
}

func (uc *userController) FindUser(c *gin.Context) {
	autHeader := c.GetHeader("Authorization")
	token, err := uc.jwtService.ValidateToken(autHeader)
	if err != nil {
		panic(err.Error())
	}
	claim := token.Claims.(jwt.MapClaims)
	usr := uc.userService.FindUser(fmt.Sprintf("%v", claim["user_id"]))
	res := helper.BuildResponse(true, "OK", usr)
	c.JSON(http.StatusOK, res)
}
