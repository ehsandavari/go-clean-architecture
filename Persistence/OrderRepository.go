package Persistence

import (
	"GolangCodeBase/Application/Common/Interfaces"
	"GolangCodeBase/Domain/Entities"
)

type sOrderRepository struct {
	Interfaces.IGenericRepository[Entities.OrderEntity]
}

func newOrderRepository(db *sDatabaseContext) Interfaces.IOrderRepository {
	return sOrderRepository{
		IGenericRepository: newGenericRepository[Entities.OrderEntity](db),
	}
}
