package persistence

import (
	"github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/golang-clean-architecture/domain/entities"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/postgres/models"
)

type sOrderRepository struct {
	interfaces.IGenericRepository[entities.OrderEntity]
}

func newOrderRepository(db *SDatabaseContext) interfaces.IOrderRepository {
	return sOrderRepository{
		IGenericRepository: newGenericRepository[models.OrderModel, entities.OrderEntity](db),
	}
}
