package persistence

import (
	"github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/golang-clean-architecture/domain/entities"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/postgres/models"
	"github.com/google/uuid"
)

type sGenericRepository[TD models.IModel[TE], TE entities.IEntityConstraint] struct {
	dataBaseContext *SDatabaseContext
}

func newGenericRepository[TD models.IModel[TE], TE entities.IEntityConstraint](dataBaseContext *SDatabaseContext) interfaces.IGenericRepository[TE] {
	return sGenericRepository[TD, TE]{
		dataBaseContext: dataBaseContext,
	}
}

func (r sGenericRepository[TD, TE]) First() TE {
	var model TD
	r.dataBaseContext.Postgres.DB.First(&model)
	return model.ToEntity()
}

func (r sGenericRepository[TD, TE]) Last() TE {
	var model TD
	r.dataBaseContext.Postgres.DB.Last(&model)
	return model.ToEntity()
}

func (r sGenericRepository[TD, TE]) All() []TE {
	var modelsObject []TD
	r.dataBaseContext.Postgres.DB.Find(&modelsObject)
	var entitiesObject []TE
	for _, model := range modelsObject {
		entitiesObject = append(entitiesObject, model.ToEntity())
	}
	return entitiesObject
}

func (r sGenericRepository[TD, TE]) Add(entity TE) int64 {
	var model TD
	model = model.FromEntity(entity).(TD)
	return r.dataBaseContext.Postgres.DB.Create(&model).RowsAffected
}

func (r sGenericRepository[TD, TE]) Update(id uuid.UUID, entity TE) int64 {
	var model TD
	model = model.FromEntity(entity).(TD)
	return r.dataBaseContext.Postgres.DB.Where("id", id).Updates(&model).RowsAffected
}

func (r sGenericRepository[TD, TE]) Delete(id uuid.UUID) int64 {
	var model TD
	return r.dataBaseContext.Postgres.DB.Delete(&model, id).RowsAffected
}
