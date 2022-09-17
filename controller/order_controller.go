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

type OrderController interface {
	All(c *gin.Context)
	Insert(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type orderController struct {
	orderService service.OrderService
}

func NewOrderController(orderServ service.OrderService) OrderController {
	return &orderController{
		orderService: orderServ,
	}
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
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result := oc.orderService.Insert(orderCreateDTO)
	response := helper.BuildResponse(true, "OK", result)
	c.JSON(http.StatusCreated, response)
	return
}

func (oc *orderController) Update(c *gin.Context) {

	var orderDTO dto.OrderUpdateDTO

	if err := c.ShouldBind(&orderDTO); err != nil {
		res := helper.BuildErrorResponse("Failed to bind order", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := oc.orderService.Update(orderDTO)
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
