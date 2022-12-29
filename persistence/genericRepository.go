package persistence

import (
	"golangCodeBase/application/common/interfaces"
	"golangCodeBase/domain/entities"
)

type sGenericRepository[T entities.IEntityConstraint] struct {
	dataBaseContext *sDatabaseContext
}

func newGenericRepository[T entities.IEntityConstraint](dataBaseContext *sDatabaseContext) interfaces.IGenericRepository[T] {
	return sGenericRepository[T]{
		dataBaseContext: dataBaseContext,
	}
}

func (r sGenericRepository[T]) Find() (find T) {
	r.dataBaseContext.postgres.DB.First(&find)
	return find
}

func (r sGenericRepository[T]) Add(model T) T {
	r.dataBaseContext.postgres.DB.Create(&model)
	return model
}
