package Commands

import (
	"GolangCodeBase/Application/Common"
	ApplicationInterfaces "GolangCodeBase/Application/Common/Interfaces"
	"GolangCodeBase/Domain/Interfaces"
	DomainInterfaces "GolangCodeBase/Domain/Interfaces"
	"GolangCodeBase/Infrastructure/Config"
	"context"
	"fmt"
)

type SPublishOrderCommand struct {
	Price uint   `json:"price" validate:"required,gte=0,email"`
	Title string `json:"title" validate:"required,gte=0"`
}

type sPublishOrderCommandHandler struct {
	sConfig     *Config.SConfig
	iLogger     ApplicationInterfaces.ILogger
	iUnitOfWork DomainInterfaces.IUnitOfWork
	iRedis      ApplicationInterfaces.IRedis
}

func newPublishOrderCommandHandler(
	sConfig *Config.SConfig,
	iLogger ApplicationInterfaces.ILogger,
	iUnitOfWork DomainInterfaces.IUnitOfWork,
	iRedis ApplicationInterfaces.IRedis,
) ApplicationInterfaces.IBaseCommand[SPublishOrderCommand] {
	return sPublishOrderCommandHandler{
		sConfig:     sConfig,
		iLogger:     iLogger,
		iUnitOfWork: iUnitOfWork,
		iRedis:      iRedis,
	}
}

func (r *sPublishOrderCommandHandler) Handle(ctx context.Context, command SPublishOrderCommand) {
	err := r.iUnitOfWork.Do(func(work Interfaces.IUnitOfWork) error {
		work.OrderRepository().Add(command)
		return nil
	})
	fmt.Println(err)
	err = r.iRedis.Publish(ctx, r.sConfig.Redis.Queues["Orders"], Common.MarshalJson(command))
	if err != nil {
		panic(err)
	}
}
