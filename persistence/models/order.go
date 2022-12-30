package models

import (
	"golangCodeBase/domain/entities"
)

type OrderModel struct {
	Id    uint64 `gorm:"primary_key"`
	Price uint
	Title string
}

func (OrderModel) TableName() string {
	return "orders"
}

func (r OrderModel) ToEntity() entities.OrderEntity {
	return entities.OrderEntity{
		Price: r.Price,
		Title: r.Title,
	}
}

func (r OrderModel) FromEntity(entity entities.OrderEntity) any {
	r.Price = entity.Price
	r.Title = entity.Title
	return r
}
