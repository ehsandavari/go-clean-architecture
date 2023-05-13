package createOrder

import (
	"context"
	"github.com/ehsandavari/go-mediator"
	"github.com/ehsandavari/golang-clean-architecture/application"
	"github.com/ehsandavari/golang-clean-architecture/application/common"
	"github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/golang-clean-architecture/domain/entities"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/config"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"time"
)

func init() {
	application.Modules = append(application.Modules, fx.Invoke(func(
		sConfig *config.SConfig,
		iLogger interfaces.ILogger,
		iUnitOfWork interfaces.IUnitOfWork,
	) {
		if err := mediator.RegisterRequestHandler[SCreateOrderCommand, *entities.Order](newCreateOrderCommandHandler(sConfig, iLogger, iUnitOfWork)); err != nil {
			iLogger.Fatal(err)
		}
	}))
}

type SCreateOrderCommandHandler struct {
	sConfig     *config.SConfig
	iLogger     interfaces.ILogger
	iUnitOfWork interfaces.IUnitOfWork
}

func newCreateOrderCommandHandler(
	sConfig *config.SConfig,
	iLogger interfaces.ILogger,
	iUnitOfWork interfaces.IUnitOfWork,
) SCreateOrderCommandHandler {
	return SCreateOrderCommandHandler{
		sConfig:     sConfig,
		iLogger:     iLogger,
		iUnitOfWork: iUnitOfWork,
	}
}

func (r SCreateOrderCommandHandler) Handle(ctx context.Context, command SCreateOrderCommand) (*entities.Order, mediator.IError) {
	orderEntity := entities.NewOrder(
		uuid.New(),
		command.UserId,
		command.Title,
		command.Price,
		entities.NewBase(time.Now(), time.Now()),
	)
	createdOrderNumber, err := r.iUnitOfWork.OrderRepository().Create(ctx, orderEntity)
	if err != nil {
		return nil, common.ErrorInternal
	}
	if createdOrderNumber == 0 {
		r.iLogger.Warn("order not created")
		return nil, common.ErrorInternal
	}
	return orderEntity, nil
}
