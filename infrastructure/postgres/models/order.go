package models

import (
	"github.com/ehsandavari/golang-clean-architecture/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderModel struct {
	BaseModel[uuid.UUID]
	Price uint
	Title string
}

func (OrderModel) TableName() string {
	return "orders"
}

func (r *OrderModel) BeforeCreate(*gorm.DB) error {
	r.ID = uuid.New()
	return nil
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
