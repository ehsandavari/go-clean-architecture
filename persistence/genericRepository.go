package persistence

import (
	"github.com/google/uuid"
	"golangCodeBase/application/common/interfaces"
	"golangCodeBase/domain/entities"
)

type sGenericRepository[TD IDataModel[TE], TE entities.IEntityConstraint] struct {
	dataBaseContext *SDatabaseContext
}

func newGenericRepository[TD IDataModel[TE], TE entities.IEntityConstraint](dataBaseContext *SDatabaseContext) interfaces.IGenericRepository[TE] {
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
	var models []TD
	r.dataBaseContext.Postgres.DB.Find(&models)
	var entitiesObject []TE
	for _, model := range models {
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
