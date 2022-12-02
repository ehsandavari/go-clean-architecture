package Interfaces

import "GolangCodeBase/Domain/Entities"

//go:generate mockgen -destination=../Mocks/MockIGenericRepository.go -package=Mock GolangCodeBase/Domain/Interfaces IGenericRepository

type GenericRepositoryConstraint interface {
	Entities.OrderEntity
}

type IGenericRepository[T GenericRepositoryConstraint] interface {
	Find() T
	Add(model T) T
}
