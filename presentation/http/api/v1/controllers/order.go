package controllers

import (
	"github.com/ehsandavari/go-mediator"
	"github.com/ehsandavari/golang-clean-architecture/application/common"
	"github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/golang-clean-architecture/application/handlers/order/queries/GetAllOrderByFilter"
	"github.com/ehsandavari/golang-clean-architecture/domain/entities"
	"github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"net/http"
)

type sOrderController struct {
	sBaseController
}

func NewOrderController(routerGroup *gin.RouterGroup, logger interfaces.ILogger) {
	orderController := sOrderController{
		sBaseController: newBaseController(logger),
	}

	routerGroup = routerGroup.Group("/order")
	{
		routerGroup.POST("/all", baseController[*common.PaginateResult[entities.Order]](orderController.all).Handle())
	}
}

// GetOrders
//
//	@Tags			order
//	@Summary		Get All Order
//	@ID				get-map
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Language	header		string				false	"some description"	Enums(en, fa)
//	@Param			params			body		GetAllOrderRequest	false	"Query Params"
//	@Success		200				{object}	BaseApiResponse[common.PaginateResult[entities.Order]]
//	@Router			/order/all [POST]
func (r *sOrderController) all(ctx *gin.Context) BaseApiResponse[*common.PaginateResult[entities.Order]] {
	var paginateQuery common.PaginateQuery
	if err := ctx.ShouldBindJSON(&paginateQuery); err != nil {
		return newBaseApiResponse[*common.PaginateResult[entities.Order]](
			err,
			err.Error(),
			1,
			nil,
			http.StatusBadRequest,
		)
	}
	r.iLogger.Info("paginateQuery: ", paginateQuery)

	list, err := mediator.Send[GetAllOrderByFilter.SGetAllOrderByFilterQuery, *common.PaginateResult[entities.Order]](ctx, GetAllOrderByFilter.NewSGetAllOrderByFilterQuery(paginateQuery))
	if err != nil {
		return newBaseApiResponse[*common.PaginateResult[entities.Order]](
			err,
			i18n.MustGetMessage(err.Error()),
			err.Code(),
			nil,
			http.StatusBadRequest,
		)
	}

	return newBaseApiResponse[*common.PaginateResult[entities.Order]](
		nil,
		"success",
		200,
		list,
		http.StatusOK,
	)
}
