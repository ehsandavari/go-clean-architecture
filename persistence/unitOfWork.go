package persistence

import (
	"go.uber.org/fx"
	"golangCodeBase/application/common/interfaces"
	"golangCodeBase/infrastructure/postgres"
)

func init() {
	Modules = append(Modules, fx.Provide(NewUnitOfWork))
}

type sUnitOfWork struct {
	databaseContext *sDatabaseContext
	orderRepository interfaces.IOrderRepository
}

func NewUnitOfWork(databaseContext *sDatabaseContext) interfaces.IUnitOfWork {
	return &sUnitOfWork{
		databaseContext: databaseContext,
		orderRepository: newOrderRepository(databaseContext),
	}
}

func (r sUnitOfWork) OrderRepository() interfaces.IOrderRepository {
	return r.orderRepository
}

func (r sUnitOfWork) Do(unitOfWorkBlock interfaces.UnitOfWorkBlock) error {
	return r.databaseContext.postgres.Transaction(func(transaction *postgres.SPostgres) error {
		r.databaseContext.postgres = transaction
		r.orderRepository = newOrderRepository(r.databaseContext)
		return unitOfWorkBlock(r)
	})
}
