package interfaces

import (
	"github.com/ehsandavari/golang-clean-architecture/application/common"
	"github.com/ehsandavari/golang-clean-architecture/domain/entities"
	"github.com/google/uuid"
)

//go:generate mockgen -destination=../../mocks/mockPersistence.go -package=mocks golangCodeBase/application/common/interfaces IUnitOfWork,IOrderRepository

type (
	IGenericRepository[TE entities.IEntityConstraint] interface {
		Paginate(listQuery common.PaginateQuery) (*common.PaginateResult[TE], error)
		All() []TE
		First() TE
		Last() TE
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
