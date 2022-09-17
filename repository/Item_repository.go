package repository

import (
	"assignment2go/entity"

	"gorm.io/gorm"
)

type ItemRepository interface {
	InsertItem(item entity.Item) (entity.Item, error)
}

type itemConnection struct {
	connection *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemConnection{
		connection: db,
	}
}

func (db *itemConnection) InsertItem(item entity.Item) (entity.Item, error) {
	if err := db.connection.Create(&item).Error; err != nil {
		return item, err
	}
	return item, nil
}
