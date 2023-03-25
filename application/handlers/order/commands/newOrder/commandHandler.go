package publishOrder

import (
	"context"
	"github.com/ehsandavari/go-mediator"
	"github.com/ehsandavari/golang-clean-architecture/application"
	"github.com/ehsandavari/golang-clean-architecture/application/common"
	ApplicationInterfaces "github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/config"
	"go.uber.org/fx"
)

func init() {
	application.Modules = append(application.Modules, fx.Invoke(func(
		sConfig *config.SConfig,
		iLogger ApplicationInterfaces.ILogger,
		iUnitOfWork ApplicationInterfaces.IUnitOfWork,
		iRedis ApplicationInterfaces.IRedis,
	) {
		if err := mediator.RegisterRequestHandler[SPublishOrderCommand, string](
			newPublishOrderCommandHandler(sConfig, iLogger, iUnitOfWork, iRedis),
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

func newPublishOrderCommandHandler(
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

func (r SPublishOrderCommandHandler) Handle(ctx context.Context, command SPublishOrderCommand) (string, mediator.IError) {
	err := r.iRedis.Publish(ctx, r.sConfig.Redis.Queues["Orders"], common.MarshalJson(command))
	if err != nil {
		return "", common.ErrorInternal
	}
	return "", nil
}
