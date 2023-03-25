package GetAllOrderByFilter

import (
	"context"
	"github.com/ehsandavari/go-mediator"
	"github.com/ehsandavari/golang-clean-architecture/application"
	"github.com/ehsandavari/golang-clean-architecture/application/common"
	ApplicationInterfaces "github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/golang-clean-architecture/domain/entities"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/config"
	"go.uber.org/fx"
)

func init() {
	application.Modules = append(application.Modules, fx.Invoke(func(
		sConfig *config.SConfig,
		iLogger ApplicationInterfaces.ILogger,
		iUnitOfWork ApplicationInterfaces.IUnitOfWork,
	) {
		if err := mediator.RegisterRequestHandler[SGetAllOrderByFilterQuery, *common.PaginateResult[entities.OrderEntity]](
			newGetAllOrderByFilterQueryHandler(sConfig, iLogger, iUnitOfWork),
		); err != nil {
			panic(err)
		}
	}))
}

type SGetAllOrderByFilterQueryHandler struct {
	sConfig     *config.SConfig
	iLogger     ApplicationInterfaces.ILogger
	iUnitOfWork ApplicationInterfaces.IUnitOfWork
}

func newGetAllOrderByFilterQueryHandler(
	sConfig *config.SConfig,
	iLogger ApplicationInterfaces.ILogger,
	iUnitOfWork ApplicationInterfaces.IUnitOfWork,
) SGetAllOrderByFilterQueryHandler {
	return SGetAllOrderByFilterQueryHandler{
		sConfig:     sConfig,
		iLogger:     iLogger,
		iUnitOfWork: iUnitOfWork,
	}
}

func (r SGetAllOrderByFilterQueryHandler) Handle(ctx context.Context, Query SGetAllOrderByFilterQuery) (*common.PaginateResult[entities.OrderEntity], error) {
	paginate, err := r.iUnitOfWork.OrderRepository().Paginate(Query.PaginateQuery)
	if err != nil {
		return nil, common.ErrorOrderNotFound
	}
	return paginate, nil
}
