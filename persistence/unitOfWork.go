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
	databaseContext *SDatabaseContext
	orderRepository interfaces.IOrderRepository
}

func NewUnitOfWork(databaseContext *SDatabaseContext) interfaces.IUnitOfWork {
	return &sUnitOfWork{
		databaseContext: databaseContext,
		orderRepository: newOrderRepository(databaseContext),
	}
}

func (r sUnitOfWork) OrderRepository() interfaces.IOrderRepository {
	return r.orderRepository
}

func (r sUnitOfWork) Do(unitOfWorkBlock func(interfaces.IUnitOfWork) error) error {
	return r.databaseContext.Postgres.Transaction(func(transaction *postgres.SPostgres) error {
		r.databaseContext.Postgres = transaction
		r.orderRepository = newOrderRepository(r.databaseContext)
		return unitOfWorkBlock(r)
	})
}
