package Interfaces

import "GolangCodeBase/Domain/Entities"

//go:generate mockgen -destination=../Mocks/MockIGenericRepository.go -package=Mock GolangCodeBase/Domain/Interfaces IGenericRepository

type IGenericRepository[T Entities.IEntityConstraint] interface {
	Find() T
	Add(model T) T
}
