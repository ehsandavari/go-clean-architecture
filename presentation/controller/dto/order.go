package dto

import DomainEntities "github.com/ehsandavari/golang-clean-architecture/domain/entities"

type CreateOrderRequest struct {
	Id    uint64 `json:"order_id"`
	Price uint   `json:"price"`
	Title string `json:"title"`
}

type CreateOrderResponse struct {
	User DomainEntities.OrderEntity `json:"user"`
}
