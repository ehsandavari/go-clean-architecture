package models

import (
	"github.com/ehsandavari/golang-clean-architecture/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID    uuid.UUID `gorm:"type:uuid;"`
	Price uint
	Title string
	BaseModel
}

func (Order) TableName() string {
	return "orders"
}

func (r *Order) BeforeCreate(*gorm.DB) error {
	r.ID = uuid.New()
	return nil
}

func (r Order) ToEntity() entities.Order {
	return entities.Order{
		Id:    r.ID,
		Price: r.Price,
		Title: r.Title,
	}
}

func (r Order) FromEntity(entity entities.Order) any {
	r.ID = entity.Id
	r.Price = entity.Price
	r.Title = entity.Title
	return r
}
