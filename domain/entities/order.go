package entities

import "github.com/google/uuid"

type Order struct {
	Id    uuid.UUID
	Price uint
	Title string
} //@name Order

func NewOrder(price uint, title string) Order {
	return Order{Price: price, Title: title}
}
