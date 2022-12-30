package persistence

import (
	"golangCodeBase/application/common/interfaces"
	"golangCodeBase/domain/entities"
	"golangCodeBase/persistence/models"
)

type sGenericRepository[TD models.DataModel[TE], TE entities.IEntityConstraint] struct {
	dataBaseContext *SDatabaseContext
}

func newGenericRepository[TD models.DataModel[TE], TE entities.IEntityConstraint](dataBaseContext *SDatabaseContext) interfaces.IGenericRepository[TE] {
	return sGenericRepository[TD, TE]{
		dataBaseContext: dataBaseContext,
	}
}

func (r sGenericRepository[TD, TE]) Find() TE {
	var model TD
	r.dataBaseContext.Postgres.DB.First(&model)
	return model.ToEntity()
}

func (r sGenericRepository[TD, TE]) Add(entity TE) int64 {
	var model TD
	model = model.FromEntity(entity).(TD)
	return r.dataBaseContext.Postgres.DB.Create(&model).RowsAffected
}
