package models

import (
	"github.com/ehsandavari/golang-clean-architecture/domain/entities"
)

type OrderModel struct {
	Base
	Price uint
	Title string
}

func (OrderModel) TableName() string {
	return "orders"
}

func (r OrderModel) ToEntity() entities.OrderEntity {
	return entities.OrderEntity{
		Id:    r.ID,
		Price: r.Price,
		Title: r.Title,
	}
}

func (r OrderModel) FromEntity(entity entities.OrderEntity) any {
	r.ID = entity.Id
	r.Price = entity.Price
	r.Title = entity.Title
	return r
}
