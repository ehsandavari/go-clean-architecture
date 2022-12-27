package Persistence

import (
	"GolangCodeBase/Application/Common/Interfaces"
	"GolangCodeBase/Domain/Entities"
)

type sGenericRepository[T Entities.IEntityConstraint] struct {
	dataBaseContext *sDatabaseContext
}

func newGenericRepository[T Entities.IEntityConstraint](dataBaseContext *sDatabaseContext) Interfaces.IGenericRepository[T] {
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
