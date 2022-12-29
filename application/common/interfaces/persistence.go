package interfaces

import (
	"golangCodeBase/domain/entities"
)

//go:generate mockgen -destination=../../mocks/mockPersistence.go -package=mocks golangCodeBase/application/common/interfaces IUnitOfWork,IOrderRepository

type (
	IGenericRepository[T entities.IEntityConstraint] interface {
		Find() T
		Add(model T) T
	}
	IOrderRepository interface {
		IGenericRepository[entities.OrderEntity]
	}
	UnitOfWorkBlock func(IUnitOfWork) error
	IUnitOfWork     interface {
		OrderRepository() IOrderRepository
		Do(UnitOfWorkBlock) error
	}
)
