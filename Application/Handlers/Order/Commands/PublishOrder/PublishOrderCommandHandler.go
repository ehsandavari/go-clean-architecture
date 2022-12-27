package PublishOrder

import (
	"GolangCodeBase/Application"
	"GolangCodeBase/Application/Common"
	ApplicationInterfaces "GolangCodeBase/Application/Common/Interfaces"
	"GolangCodeBase/Infrastructure/Config"
	"context"
	"github.com/mehdihadeli/go-mediatr"
	"go.uber.org/fx"
)

type SPublishOrderCommandHandler struct {
	sConfig     *Config.SConfig
	iLogger     ApplicationInterfaces.ILogger
	iUnitOfWork ApplicationInterfaces.IUnitOfWork
	iRedis      ApplicationInterfaces.IRedis
}

func NewPublishOrderCommandHandler(
	sConfig *Config.SConfig,
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

func init() {
	Application.Modules = append(Application.Modules, fx.Invoke(func(
		sConfig *Config.SConfig,
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

func (r SPublishOrderCommandHandler) Handle(ctx context.Context, command SPublishOrderCommand) (string, error) {
	err := r.iRedis.Publish(ctx, r.sConfig.Redis.Queues["Orders"], Common.MarshalJson(command))
	if err != nil {
		return "", err
	}
	return "", nil
}
