package persistence

import (
	"golangCodeBase/application/common/interfaces"
	"golangCodeBase/domain/entities"
)

type sOrderRepository struct {
	interfaces.IGenericRepository[entities.OrderEntity]
}

func newOrderRepository(db *sDatabaseContext) interfaces.IOrderRepository {
	return sOrderRepository{
		IGenericRepository: newGenericRepository[entities.OrderEntity](db),
	}
}
