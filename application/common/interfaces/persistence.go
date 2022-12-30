package interfaces

import (
	"golangCodeBase/domain/entities"
)

//go:generate mockgen -destination=../../mocks/mockPersistence.go -package=mocks golangCodeBase/application/common/interfaces IUnitOfWork,IOrderRepository

type (
	IGenericRepository[TE entities.IEntityConstraint] interface {
		Find() TE
		Add(model TE) int64
	}
	IOrderRepository interface {
		IGenericRepository[entities.OrderEntity]
	}
	IUnitOfWork interface {
		OrderRepository() IOrderRepository
		Do(func(IUnitOfWork) error) error
	}
)
