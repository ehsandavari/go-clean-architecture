package Persistence

import (
	"GolangCodeBase/Domain/Interfaces"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type sUnitOfWork struct {
	databaseContext *sDatabaseContext
	orderRepository Interfaces.IOrderRepository
}

type sUnitOfWorkParams struct {
	fx.In
	SDatabaseContext *sDatabaseContext
	IOrderRepository Interfaces.IOrderRepository
}

func NewUnitOfWork(unitOfWorkParams sUnitOfWorkParams) Interfaces.IUnitOfWork {
	return &sUnitOfWork{
		databaseContext: unitOfWorkParams.SDatabaseContext,
		orderRepository: unitOfWorkParams.IOrderRepository,
	}
}

func (r sUnitOfWork) OrderRepository() Interfaces.IOrderRepository {
	return r.orderRepository
}

func (r sUnitOfWork) Do(unitOfWorkBlock Interfaces.UnitOfWorkBlock) error {
	return r.databaseContext.postgres.DB.Transaction(func(transaction *gorm.DB) error {
		//r.databaseContext.Postgres().DB = transaction
		//r.orderRepository = newOrderRepository(r.databaseContext)
		return unitOfWorkBlock(r)
	})
}
