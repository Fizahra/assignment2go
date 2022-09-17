package repository

import (
	"assignment2go/entity"
	"context"

	"gorm.io/gorm"
)

type OrderRepository interface {
	InsertOrder(order entity.Order) entity.Order
	UpdateOrder(order entity.Order) (entity.Order, error)
	DeleteOrder(context.Context, uint64) error
	ReadOrder() []entity.Order
	FindOrderByID(orderID uint64) entity.Order
}

type orderConnection struct {
	connection *gorm.DB
}

func NewOrderRepository(dbConn *gorm.DB) OrderRepository {
	return &orderConnection{
		connection: dbConn,
	}
}

func (oc *orderConnection) InsertOrder(order entity.Order) entity.Order {
	oc.connection.Save(&order)
	oc.connection.Preload("Item").Find(&order)
	return order
}

func (oc *orderConnection) UpdateOrder(order entity.Order) (entity.Order, error) {
	if err := oc.connection.Debug().Model(&order).Association("Item").Replace(order.Item); err != nil {
		return order, err
	}
	if err := oc.connection.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Preload("Item").Save(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (oc *orderConnection) DeleteOrder(ctx context.Context, id uint64) error {
	tx := oc.connection.WithContext(ctx).Where("order_id = ?", id).Delete(&entity.Order{})

	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (oc *orderConnection) FindOrderByID(orderID uint64) entity.Order {
	var order entity.Order
	oc.connection.Preload("Item").Find(&order, orderID)
	return order
}

func (oc *orderConnection) ReadOrder() []entity.Order {
	var orders []entity.Order
	oc.connection.Preload("Item").Find(&orders)
	return orders
}
