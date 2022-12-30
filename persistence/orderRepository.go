package persistence

import (
	"golangCodeBase/application/common/interfaces"
	"golangCodeBase/domain/entities"
	"golangCodeBase/persistence/models"
)

type sOrderRepository struct {
	interfaces.IGenericRepository[entities.OrderEntity]
}

func newOrderRepository(db *SDatabaseContext) interfaces.IOrderRepository {
	return sOrderRepository{
		IGenericRepository: newGenericRepository[models.OrderModel, entities.OrderEntity](db),
	}
}
