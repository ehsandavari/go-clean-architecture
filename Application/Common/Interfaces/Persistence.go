package Interfaces

import (
	"GolangCodeBase/Domain/Entities"
)

//go:generate mockgen -destination=../../Mocks/MockPersistence.go -package=Mocks GolangCodeBase/Application/Common/Interfaces IUnitOfWork,IOrderRepository

type (
	IGenericRepository[T Entities.IEntityConstraint] interface {
		Find() T
		Add(model T) T
	}
	IOrderRepository interface {
		IGenericRepository[Entities.OrderEntity]
	}
	UnitOfWorkBlock func(IUnitOfWork) error
	IUnitOfWork     interface {
		OrderRepository() IOrderRepository
		Do(UnitOfWorkBlock) error
	}
)
