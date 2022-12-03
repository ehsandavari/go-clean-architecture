package Persistence

import (
	"GolangCodeBase/Domain/Entities"
	"GolangCodeBase/Domain/Interfaces"
)

type sOrderRepository struct {
	Interfaces.IGenericRepository[Entities.OrderEntity]
}

func NewOrderRepository(db *sDatabaseContext) Interfaces.IOrderRepository {
	return sOrderRepository{
		IGenericRepository: newGenericRepository[Entities.OrderEntity](db),
	}
}
