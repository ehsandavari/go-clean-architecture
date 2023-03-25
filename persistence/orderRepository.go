package persistence

import (
	"github.com/ehsandavari/golang-clean-architecture/domain/entities"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/postgres/models"
)

type sOrderRepository struct {
	sGenericRepository[models.Order, entities.Order]
}

func newOrderRepository(db *SDatabaseContext) sOrderRepository {
	return sOrderRepository{
		sGenericRepository: newGenericRepository[models.Order, entities.Order](db),
	}
}

func (r sOrderRepository) FindById() entities.Order {
	var entitiesObjects entities.Order
	r.Postgres.DB.Model(new(models.Order)).First(&entitiesObjects)
	return entitiesObjects
}
