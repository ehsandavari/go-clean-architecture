package subscribeOrder

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ehsandavari/go-mediator"
	"github.com/ehsandavari/golang-clean-architecture/application"
	"github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/golang-clean-architecture/domain/entities"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/config"
	"go.uber.org/fx"
)

func init() {
	application.Modules = append(application.Modules, fx.Invoke(func(
		sConfig *config.SConfig,
		iLogger interfaces.ILogger,
		iUnitOfWork interfaces.IUnitOfWork,
		iRedis interfaces.IRedis,
	) {
		if err := mediator.RegisterRequestHandler[SSubscribeOrderCommand, string](
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
		channelData := <-channel
		if err := json.Unmarshal([]byte(channelData), &orderEntity); err != nil {
			panic(err)
		}
		add := r.iUnitOfWork.OrderRepository().Add(orderEntity)
		fmt.Println(add)
	}()
	return "", nil
}
