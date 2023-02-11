package persistence

import (
	"github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/golang-clean-architecture/domain/entities"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/postgres/models"
	"github.com/google/uuid"
)

type sGenericRepository[TM models.IModel[TE], TE entities.IEntityConstraint] struct {
	dataBaseContext *SDatabaseContext
}

func newGenericRepository[TM models.IModel[TE], TE entities.IEntityConstraint](dataBaseContext *SDatabaseContext) interfaces.IGenericRepository[TE] {
	return sGenericRepository[TM, TE]{
		dataBaseContext: dataBaseContext,
	}
}

func (r sGenericRepository[TM, TE]) First() TE {
	var model TM
	r.dataBaseContext.Postgres.DB.First(&model)
	return model.ToEntity()
}

func (r sGenericRepository[TM, TE]) Last() TE {
	var model TM
	r.dataBaseContext.Postgres.DB.Last(&model)
	return model.ToEntity()
}

func (r sGenericRepository[TM, TE]) All() []TE {
	var model TM
	var entitiesObject []TE
	r.dataBaseContext.Postgres.DB.Model(model).Find(&entitiesObject)
	return entitiesObject
}

func (r sGenericRepository[TM, TE]) Add(entity TE) int64 {
	var model TM
	model = model.FromEntity(entity).(TM)
	return r.dataBaseContext.Postgres.DB.Create(&model).RowsAffected
}

func (r sGenericRepository[TM, TE]) Update(id uuid.UUID, entity TE) int64 {
	var model TM
	model = model.FromEntity(entity).(TM)
	return r.dataBaseContext.Postgres.DB.Where("id", id).Updates(&model).RowsAffected
}

func (r sGenericRepository[TM, TE]) Delete(id uuid.UUID) int64 {
	var model TM
	return r.dataBaseContext.Postgres.DB.Delete(&model, id).RowsAffected
}
