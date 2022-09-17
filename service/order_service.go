package service

import (
	"assignment2go/dto"
	"assignment2go/entity"
	"assignment2go/repository"
	"context"
	"log"

	"github.com/mashingan/smapping"
)

type OrderService interface {
	Insert(order dto.OrderCreateDTO) entity.Order
	Update(order dto.OrderUpdateDTO) (entity.Order, error)
	DeleteOrder(context.Context, uint64) error
	GetOrder() []entity.Order
	FindOrderByID(orderID uint64) entity.Order
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) OrderService {
	return &orderService{
		orderRepository: orderRepo,
	}
}

func (os *orderService) Insert(order dto.OrderCreateDTO) entity.Order {
	o := entity.Order{}
	err := smapping.FillStruct(&o, smapping.MapFields(&order))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := os.orderRepository.InsertOrder(o)
	return res
}

func (os *orderService) Update(order dto.OrderUpdateDTO) (entity.Order, error) {
	o := entity.Order{}
	err := smapping.FillStruct(&o, smapping.MapFields(&order))
	if err != nil {
		return o, err
	}
	result, err := os.orderRepository.UpdateOrder(o)
	return result, err
}

func (os *orderService) DeleteOrder(ctx context.Context, id uint64) error {
	return os.orderRepository.DeleteOrder(ctx, id)
}

func (os *orderService) GetOrder() []entity.Order {
	return os.orderRepository.ReadOrder()
}

func (os *orderService) FindOrderByID(orderID uint64) entity.Order {
	return os.orderRepository.FindOrderByID(orderID)
}
