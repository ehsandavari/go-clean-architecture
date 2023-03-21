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
		routerGroup.POST("/all", baseController[*common.PaginateResult[entities.OrderEntity]](orderController.all).Handle())
	}
}

// GetOrders
//
//	@Tag			order
//	@Summary		Get Map Order
//	@Description	get map
//	@ID				get-map
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Language	header		string				false	"some description"	Enums(en, fa)
//	@Param			params			body		GetAllOrderRequest	false	"Query Params"
//	@Success		200				{object}	BaseApiResponse[common.PaginateResult[entities.OrderEntity]]
//	@Router			/order/all [POST]
func (r *sOrderController) all(ctx *gin.Context) BaseApiResponse[*common.PaginateResult[entities.OrderEntity]] {
	var paginateQuery common.PaginateQuery
	if err := ctx.ShouldBindJSON(&paginateQuery); err != nil {
		return newBaseApiResponse[*common.PaginateResult[entities.OrderEntity]](
			err,
			err.Error(),
			22,
			nil,
			http.StatusBadRequest,
		)
	}
	r.iLogger.Info("paginateQuery: ", paginateQuery)

	list, err := mediator.Send[GetAllOrderByFilter.SGetAllOrderByFilterQuery, *common.PaginateResult[entities.OrderEntity]](ctx, GetAllOrderByFilter.NewSGetAllOrderByFilterQuery(paginateQuery))
	if err != nil {
		return newBaseApiResponse[*common.PaginateResult[entities.OrderEntity]](
			err,
			i18n.MustGetMessage("Order."+err.Error()),
			22,
			nil,
			http.StatusBadRequest,
		)
	}

	return newBaseApiResponse[*common.PaginateResult[entities.OrderEntity]](
		nil,
		"success",
		200,
		list,
		http.StatusOK,
	)
}
