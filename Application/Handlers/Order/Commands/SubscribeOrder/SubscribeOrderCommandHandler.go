package SubscribeOrder

import (
	"GolangCodeBase/Application"
	"GolangCodeBase/Application/Common/Interfaces"
	"GolangCodeBase/Domain/Entities"
	"GolangCodeBase/Infrastructure/Config"
	"context"
	"encoding/json"
	"fmt"
	"github.com/mehdihadeli/go-mediatr"
	"go.uber.org/fx"
)

type SSubscribeOrderCommandHandler struct {
	sConfig     *Config.SConfig
	iLogger     Interfaces.ILogger
	iUnitOfWork Interfaces.IUnitOfWork
	iRedis      Interfaces.IRedis
}

func NewSubscribeOrderCommandHandler(
	sConfig *Config.SConfig,
	iLogger Interfaces.ILogger,
	iUnitOfWork Interfaces.IUnitOfWork,
	iRedis Interfaces.IRedis,
) SSubscribeOrderCommandHandler {
	return SSubscribeOrderCommandHandler{
		sConfig:     sConfig,
		iLogger:     iLogger,
		iUnitOfWork: iUnitOfWork,
		iRedis:      iRedis,
	}
}

func init() {
	Application.Modules = append(Application.Modules, fx.Invoke(func(
		sConfig *Config.SConfig,
		iLogger Interfaces.ILogger,
		iUnitOfWork Interfaces.IUnitOfWork,
		iRedis Interfaces.IRedis,
	) {
		if err := mediatr.RegisterRequestHandler[SSubscribeOrderCommand, string](
			NewSubscribeOrderCommandHandler(sConfig, iLogger, iUnitOfWork, iRedis),
		); err != nil {
			panic(err)
		}
	}))
}

func (r SSubscribeOrderCommandHandler) Handle(ctx context.Context, command SSubscribeOrderCommand) (string, error) {
	channel := r.iRedis.Subscribe(ctx, r.sConfig.Redis.Queues["Orders"])
	go func() {
		orderEntity := Entities.OrderEntity{}
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
