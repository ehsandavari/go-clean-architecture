package entities

import "github.com/google/uuid"

type OrderEntity struct {
	Id    uuid.UUID
	Price uint
	Title string
} //@name OrderEntity

func NewOrderEntity(price uint, title string) OrderEntity {
	return OrderEntity{Price: price, Title: title}
}
