package persistence

import (
	"github.com/ehsandavari/golang-clean-architecture/domain/entities"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/postgres/models"
)

type sOrderRepository struct {
	sGenericRepository[models.OrderModel, entities.OrderEntity]
}

func newOrderRepository(db *SDatabaseContext) sOrderRepository {
	return sOrderRepository{
		sGenericRepository: newGenericRepository[models.OrderModel, entities.OrderEntity](db),
	}
}

func (r sOrderRepository) FindById() entities.OrderEntity {
	var entitiesObjects entities.OrderEntity
	r.Postgres.DB.Model(new(models.OrderModel)).First(&entitiesObjects)
	return entitiesObjects
}
