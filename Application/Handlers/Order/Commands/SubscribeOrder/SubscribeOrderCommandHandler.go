package SubscribeOrder

import (
	"GolangCodeBase/Application"
	ApplicationInterfaces "GolangCodeBase/Application/Common/Interfaces"
	DomainEntities "GolangCodeBase/Domain/Entities"
	DomainInterfaces "GolangCodeBase/Domain/Interfaces"
	"GolangCodeBase/Infrastructure/Config"
	"context"
	"encoding/json"
	"fmt"
	"github.com/mehdihadeli/go-mediatr"
	"go.uber.org/fx"
)

type SSubscribeOrderCommandHandler struct {
	sConfig     *Config.SConfig
	iLogger     ApplicationInterfaces.ILogger
	iUnitOfWork DomainInterfaces.IUnitOfWork
	iRedis      ApplicationInterfaces.IRedis
}

func NewSubscribeOrderCommandHandler(
	sConfig *Config.SConfig,
	iLogger ApplicationInterfaces.ILogger,
	iUnitOfWork DomainInterfaces.IUnitOfWork,
	iRedis ApplicationInterfaces.IRedis,
) SSubscribeOrderCommandHandler {
	return SSubscribeOrderCommandHandler{
		sConfig:     sConfig,
		iLogger:     iLogger,
		iUnitOfWork: iUnitOfWork,
		iRedis:      iRedis,
	}
}

func init() {
	Application.Modules = append(Application.Modules, fx.Invoke(registerHandler))
}

func registerHandler(
	sConfig *Config.SConfig,
	iLogger ApplicationInterfaces.ILogger,
	iUnitOfWork DomainInterfaces.IUnitOfWork,
	iRedis ApplicationInterfaces.IRedis,
) {
	if err := mediatr.RegisterRequestHandler[SSubscribeOrderCommand, string](
		NewSubscribeOrderCommandHandler(sConfig, iLogger, iUnitOfWork, iRedis),
	); err != nil {
		panic(err)
	}
}

func (r SSubscribeOrderCommandHandler) Handle(ctx context.Context, command SSubscribeOrderCommand) (string, error) {
	channel := r.iRedis.Subscribe(ctx, r.sConfig.Redis.Queues["Orders"])
	go func() {
		orderEntity := DomainEntities.OrderEntity{}
		for {
			select {
			case channelData := <-channel:
				if err := json.Unmarshal([]byte(channelData), &orderEntity); err != nil {
					panic(err)
				}
				add := r.iUnitOfWork.OrderRepository().Add(orderEntity)
				fmt.Println(add)
			}
		}
	}()
	return "", nil
}
