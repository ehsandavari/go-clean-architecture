package publishOrder

import (
	"context"
	"errors"
	"github.com/ehsandavari/go-mediator"
	"github.com/ehsandavari/golang-clean-architecture/application"
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
	//f := r.iUnitOfWork.OrderRepository().First()
	//fmt.Println(f)
	//l := r.iUnitOfWork.OrderRepository().Last()
	//fmt.Println(l)
	//all := r.iUnitOfWork.OrderRepository().All()
	//fmt.Println(all)
	//a := r.iUnitOfWork.OrderRepository().Add(entities.NewOrderEntity(command.Price, command.Title))
	//fmt.Println(a)
	//u := r.iUnitOfWork.OrderRepository().Update(f.Id, entities.NewOrderEntity(command.Price+1, command.Title))
	//fmt.Println(u)
	//d := r.iUnitOfWork.OrderRepository().Delete(f.Id)
	//fmt.Println(d)
	//err := r.iRedis.Publish(ctx, r.sConfig.Redis.Queues["Orders"], common.MarshalJson(command))
	//if err != nil {
	//	return "", err
	//}
	return "", errors.New("asd")
}
