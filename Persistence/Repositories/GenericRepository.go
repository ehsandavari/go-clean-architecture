package Repositories

import (
	DomainInterfaces "GolangCodeBase/Domain/Interfaces"
)

type GenericRepository[T DomainInterfaces.GenericRepositoryConstraint] struct {
	DataBase *SDatabaseContext
}

func NewGenericRepository[T DomainInterfaces.GenericRepositoryConstraint](db *SDatabaseContext) GenericRepository[T] {
	return GenericRepository[T]{
		DataBase: db,
	}
}

func (GR GenericRepository[T]) Find() T {
	var find T
	GR.DataBase.Postgres.DB.First(&find)
	return find
}

func (GR GenericRepository[T]) Add(model T) T {
	GR.DataBase.Postgres.DB.Create(&model)
	return model
}
