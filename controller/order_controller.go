package controller

import (
	"assignment2go/dto"
	"assignment2go/entity"
	"assignment2go/helper"
	"assignment2go/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type OrderController interface {
	All(c *gin.Context)
	Insert(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type orderController struct {
	orderService service.OrderService
	jwtService   service.JWTService
}

func NewOrderController(orderServ service.OrderService, jwtServ service.JWTService) OrderController {
	return &orderController{
		orderService: orderServ,
		jwtService:   jwtServ,
	}
}

func (oc *orderController) getUserIDByToken(token string) string {
	t, err := oc.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claim := t.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%v", claim["user_id"])
}

func (oc *orderController) All(c *gin.Context) {
	var orders []entity.Order = oc.orderService.GetOrder()
	res := helper.BuildResponse(true, "OK", orders)
	c.JSON(http.StatusOK, res)
}

func (oc *orderController) Insert(c *gin.Context) {
	var orderCreateDTO dto.OrderCreateDTO
	errDTO := c.ShouldBind(&orderCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to bind order", errDTO.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	} else {
		autHeader := c.GetHeader("Authorization")
		userID := oc.getUserIDByToken(autHeader)
		convuserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			orderCreateDTO.UserID = convuserID
		}
		res := oc.orderService.Insert(orderCreateDTO)
		response := helper.BuildResponse(true, "OK", res)
		c.JSON(http.StatusCreated, response)
		return
	}
}

func (oc *orderController) Update(c *gin.Context) {
	var order dto.OrderUpdateDTO
	if err := c.ShouldBind(&order); err != nil {
		res := helper.BuildErrorResponse("Failed to bind order", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}
	result, err := oc.orderService.Update(order)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to update order", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}
	response := helper.BuildResponse(true, "OK", result)
	c.JSON(http.StatusOK, response)
	return
}

func (oc *orderController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get id", "No id found", helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if err := oc.orderService.DeleteOrder(c.Request.Context(), id); err != nil {
		res := helper.BuildErrorResponse("Failed to delete order", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}
	response := helper.BuildResponse(true, "OK", helper.EmptyObj{})
	c.JSON(http.StatusOK, response)
}
