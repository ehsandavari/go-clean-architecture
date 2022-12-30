package publishOrder

import (
	"context"
	"fmt"
	"github.com/mehdihadeli/go-mediatr"
	"go.uber.org/fx"
	"golangCodeBase/application"
	"golangCodeBase/application/common"
	ApplicationInterfaces "golangCodeBase/application/common/interfaces"
	"golangCodeBase/domain/entities"
	"golangCodeBase/infrastructure/config"
)

func init() {
	application.Modules = append(application.Modules, fx.Invoke(func(
		sConfig *config.SConfig,
		iLogger ApplicationInterfaces.ILogger,
		iUnitOfWork ApplicationInterfaces.IUnitOfWork,
		iRedis ApplicationInterfaces.IRedis,
	) {
		if err := mediatr.RegisterRequestHandler[SPublishOrderCommand, string](
			NewPublishOrderCommandHandler(sConfig, iLogger, iUnitOfWork, iRedis),
		); err != nil {
			panic(err)
		}
	}))
}

type SPublishOrderCommandHandler struct {
	sConfig     *config.SConfig
	iLogger     ApplicationInterfaces.ILogger
	iUnitOfWork ApplicationInterfaces.IUnitOfWork
	iRedis      ApplicationInterfaces.IRedis
}

func NewPublishOrderCommandHandler(
	sConfig *config.SConfig,
	iLogger ApplicationInterfaces.ILogger,
	iUnitOfWork ApplicationInterfaces.IUnitOfWork,
	iRedis ApplicationInterfaces.IRedis,
) SPublishOrderCommandHandler {
	return SPublishOrderCommandHandler{
		sConfig:     sConfig,
		iLogger:     iLogger,
		iUnitOfWork: iUnitOfWork,
		iRedis:      iRedis,
	}
}

func (r SPublishOrderCommandHandler) Handle(ctx context.Context, command SPublishOrderCommand) (string, error) {
	f := r.iUnitOfWork.OrderRepository().First()
	fmt.Println(f)
	l := r.iUnitOfWork.OrderRepository().Last()
	fmt.Println(l)
	all := r.iUnitOfWork.OrderRepository().All()
	fmt.Println(all)
	a := r.iUnitOfWork.OrderRepository().Add(entities.NewOrderEntity(command.Price, command.Title))
	fmt.Println(a)
	u := r.iUnitOfWork.OrderRepository().Update(f.Id, entities.NewOrderEntity(command.Price+1, command.Title))
	fmt.Println(u)
	d := r.iUnitOfWork.OrderRepository().Delete(f.Id)
	fmt.Println(d)
	err := r.iRedis.Publish(ctx, r.sConfig.Redis.Queues["Orders"], common.MarshalJson(command))
	if err != nil {
		return "", err
	}
	return "", nil
}