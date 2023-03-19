package dtos

import "github.com/ehsandavari/golang-clean-architecture/application/common"

type CreateOrderRequest struct {
	Id    uint64 `json:"order_id"`
	Price uint   `json:"price"`
	Title string `json:"title"`
}

type GetAllOrderRequest struct {
	common.PaginateQuery
} //@name GetAllOrderRequest
