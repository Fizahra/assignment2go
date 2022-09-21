package controller

import (
	"assignment2go/dto"
	"assignment2go/entity"
	"assignment2go/helper"
	"assignment2go/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (ac *authController) Login(c *gin.Context) {
	var login dto.LoginDTO
	err := c.ShouldBind(&login)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to login", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
	}
	authres := ac.authService.VerifyCredential(login.Email, login.Password)
	if v, ok := authres.(entity.User); ok {
		generatedtoken := ac.jwtService.GenerateToken(strconv.FormatUint(v.UserID, 10))
		v.Token = generatedtoken
		res := helper.BuildResponse(true, "Success login", v)
		c.JSON(http.StatusOK, res)
		return
	}
	res := helper.BuildErrorResponse("Oops, invalid credential!", "Make sure your credential is valid!", helper.EmptyObj{})
	c.AbortWithStatusJSON(http.StatusUnauthorized, res)
}

func (ac *authController) Register(c *gin.Context) {
	var register dto.RegisterDTO
	err := c.ShouldBind(&register)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to register!", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if !ac.authService.DuplicateEmail(register.Email) {
		res := helper.BuildErrorResponse("There's duplicat email!", "Please change your email", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusConflict, res)
		return
	} else {
		v := ac.authService.CreateUser(register)
		generatedtoken := ac.jwtService.GenerateToken(strconv.FormatUint(v.UserID, 10))
		v.Token = generatedtoken
		res := helper.BuildResponse(true, "Successfully register!", v)
		c.AbortWithStatusJSON(http.StatusCreated, res)
		return
	}
}
