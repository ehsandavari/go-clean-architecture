package Interfaces

import (
	"GolangCodeBase/Domain/Entities"
)

//go:generate mockgen -destination=../Mocks/MockIOrderRepository.go -package=Mock  GolangCodeBase/Domain/Interfaces IOrderRepository

type IOrderRepository interface {
	IGenericRepository[Entities.OrderEntity]
}
