package Repositories

import (
	DomainInterfaces "GolangCodeBase/Domain/Interfaces"
	"gorm.io/gorm"
)

type sUnitOfWork struct {
	conn            *SDatabaseContext
	orderRepository DomainInterfaces.IOrderRepository
}

func NewUnitOfWork(db *SDatabaseContext) DomainInterfaces.IUnitOfWork {
	return &sUnitOfWork{
		conn:            db,
		orderRepository: NewOrderRepository(db),
	}
}

func (r sUnitOfWork) OrderRepository() DomainInterfaces.IOrderRepository {
	return r.orderRepository
}

func (r sUnitOfWork) Do(unitOfWorkBlock DomainInterfaces.UnitOfWorkBlock) error {
	return r.conn.Postgres.DB.Transaction(func(transaction *gorm.DB) error {
		r.conn.Postgres.DB = transaction
		r.orderRepository = NewOrderRepository(r.conn)
		return unitOfWorkBlock(r)
	})
}
