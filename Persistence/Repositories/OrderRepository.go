package Repositories

import (
	DomainEntities "GolangCodeBase/Domain/Entities"
	DomainInterfaces "GolangCodeBase/Domain/Interfaces"
)

type OrderRepository struct {
	GenericRepository[DomainEntities.OrderEntity]
}

func NewOrderRepository(db *SDatabaseContext) DomainInterfaces.IOrderRepository {
	return OrderRepository{
		GenericRepository: NewGenericRepository[DomainEntities.OrderEntity](db),
	}
}
