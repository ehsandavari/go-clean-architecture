package Interfaces

//go:generate mockgen -destination=../Mocks/MockIUnitOfWork.go -package=Mock GolangCodeBase/Domain/Interfaces IUnitOfWork

type UnitOfWorkBlock func(IUnitOfWork) error
type IUnitOfWork interface {
	OrderRepository() IOrderRepository
	Do(UnitOfWorkBlock) error
}
