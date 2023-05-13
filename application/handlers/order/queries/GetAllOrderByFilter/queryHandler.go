package GetAllOrderByFilter

import (
	"context"
	"github.com/ehsandavari/go-clean-architecture/application"
	"github.com/ehsandavari/go-clean-architecture/application/common"
	"github.com/ehsandavari/go-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/go-clean-architecture/domain/entities"
	"github.com/ehsandavari/go-clean-architecture/infrastructure/config"
	"github.com/ehsandavari/go-mediator"
)

func init() {
	application.Handlers = append(application.Handlers, func(application *application.Application) {
		if err := mediator.RegisterRequestHandler[SGetAllOrderByFilterQuery, *common.PaginateResult[entities.Song]](
			newGetAllOrderByFilterQueryHandler(application.SConfig, application.ILogger, application.IUnitOfWork),
		); err != nil {
			panic(err)
		}
	})
}

type SGetAllOrderByFilterQueryHandler struct {
	sConfig     *config.SConfig
	iLogger     interfaces.ILogger
	iUnitOfWork interfaces.IUnitOfWork
}

func newGetAllOrderByFilterQueryHandler(
	sConfig *config.SConfig,
	iLogger interfaces.ILogger,
	iUnitOfWork interfaces.IUnitOfWork,
) SGetAllOrderByFilterQueryHandler {
	return SGetAllOrderByFilterQueryHandler{
		sConfig:     sConfig,
		iLogger:     iLogger,
		iUnitOfWork: iUnitOfWork,
	}
}

func (r SGetAllOrderByFilterQueryHandler) Handle(ctx context.Context, Query SGetAllOrderByFilterQuery) (*common.PaginateResult[entities.Song], mediator.IError) {
	paginate, err := r.iUnitOfWork.SongRepository().Paginate(ctx, Query.PaginateQuery)
	if err != nil {
		return nil, common.ErrorOrderNotFound
	}
	return paginate, nil
}
