package interfaces

import (
	"github.com/google/uuid"
	"golangCodeBase/domain/entities"
)

//go:generate mockgen -destination=../../mocks/mockPersistence.go -package=mocks golangCodeBase/application/common/interfaces IUnitOfWork,IOrderRepository

type (
	IGenericRepository[TE entities.IEntityConstraint] interface {
		First() TE
		Last() TE
		All() []TE
		Add(model TE) int64
		Update(id uuid.UUID, model TE) int64
		Delete(id uuid.UUID) int64
	}
	IOrderRepository interface {
		IGenericRepository[entities.OrderEntity]
	}
	IUnitOfWork interface {
		OrderRepository() IOrderRepository
		Do(func(IUnitOfWork) error) error
	}
)
