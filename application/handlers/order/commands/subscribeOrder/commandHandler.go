package subscribeOrder

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mehdihadeli/go-mediatr"
	"go.uber.org/fx"
	"golangCodeBase/application"
	"golangCodeBase/application/common/interfaces"
	"golangCodeBase/domain/entities"
	"golangCodeBase/infrastructure/config"
)

func init() {
	application.Modules = append(application.Modules, fx.Invoke(func(
		sConfig *config.SConfig,
		iLogger interfaces.ILogger,
		iUnitOfWork interfaces.IUnitOfWork,
		iRedis interfaces.IRedis,
	) {
		if err := mediatr.RegisterRequestHandler[SSubscribeOrderCommand, string](
			NewSubscribeOrderCommandHandler(sConfig, iLogger, iUnitOfWork, iRedis),
		); err != nil {
			panic(err)
		}
	}))
}

type SSubscribeOrderCommandHandler struct {
	sConfig     *config.SConfig
	iLogger     interfaces.ILogger
	iUnitOfWork interfaces.IUnitOfWork
	iRedis      interfaces.IRedis
}

func NewSubscribeOrderCommandHandler(
	sConfig *config.SConfig,
	iLogger interfaces.ILogger,
	iUnitOfWork interfaces.IUnitOfWork,
	iRedis interfaces.IRedis,
) SSubscribeOrderCommandHandler {
	return SSubscribeOrderCommandHandler{
		sConfig:     sConfig,
		iLogger:     iLogger,
		iUnitOfWork: iUnitOfWork,
		iRedis:      iRedis,
	}
}

func (r SSubscribeOrderCommandHandler) Handle(ctx context.Context, command SSubscribeOrderCommand) (string, error) {
	channel := r.iRedis.Subscribe(ctx, r.sConfig.Redis.Queues["Orders"])
	go func() {
		orderEntity := entities.OrderEntity{}
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
