package Persistence

import (
	"GolangCodeBase/Application/Common/Interfaces"
	"GolangCodeBase/Infrastructure/Postgres"
	"go.uber.org/fx"
)

func init() {
	Modules = append(Modules, fx.Provide(NewUnitOfWork))
}

type sUnitOfWork struct {
	databaseContext *sDatabaseContext
	orderRepository Interfaces.IOrderRepository
}

func NewUnitOfWork(databaseContext *sDatabaseContext) Interfaces.IUnitOfWork {
	return &sUnitOfWork{
		databaseContext: databaseContext,
		orderRepository: newOrderRepository(databaseContext),
	}
}

func (r sUnitOfWork) OrderRepository() Interfaces.IOrderRepository {
	return r.orderRepository
}

func (r sUnitOfWork) Do(unitOfWorkBlock Interfaces.UnitOfWorkBlock) error {
	return r.databaseContext.postgres.Transaction(func(transaction *Postgres.SPostgres) error {
		r.databaseContext.postgres = transaction
		r.orderRepository = newOrderRepository(r.databaseContext)
		return unitOfWorkBlock(r)
	})
}
