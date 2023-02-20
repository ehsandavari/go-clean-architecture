package dto

import "github.com/ehsandavari/golang-clean-architecture/domain/entities"

type CreateOrderRequest struct {
	Id    uint64 `json:"order_id"`
	Price uint   `json:"price"`
	Title string `json:"title"`
}

type CreateOrderResponse struct {
	User entities.OrderEntity `json:"user"`
}
