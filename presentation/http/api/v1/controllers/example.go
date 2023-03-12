package controllers

import (
	"github.com/ehsandavari/go-mediator"
	"github.com/ehsandavari/golang-clean-architecture/application/common"
	"github.com/ehsandavari/golang-clean-architecture/application/handlers/order/queries/GetAllOrderByFilter"
	"github.com/ehsandavari/golang-clean-architecture/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetMap godoc
//
//	@Tag			example
//	@Summary		Get Map Example
//	@Description	get map
//	@ID				get-map
//	@Accept			json
//	@Produce		json
//	@Param			params	body		GetAllOrderRequest	false	"Query Params"
//	@Success		200		{object}	BaseApiResponse[common.PaginateResult[entities.OrderEntity]]
//	@Failure		500		{object}	BaseApiResponse[string]
//	@Router			/ [POST]
func (r *Controller) GetMap(ctx *gin.Context) BaseApiResponse[*common.PaginateResult[entities.OrderEntity]] {
	var paginateQuery common.PaginateQuery
	if err := ctx.ShouldBindJSON(&paginateQuery); err != nil {
		return NewBaseApiResponse[*common.PaginateResult[entities.OrderEntity]](
			err,
			err.Error(),
			22,
			nil,
			http.StatusBadRequest,
		)
	}

	list, err := mediator.Send[GetAllOrderByFilter.SGetAllOrderByFilterQuery, *common.PaginateResult[entities.OrderEntity]](ctx, GetAllOrderByFilter.NewSGetAllOrderByFilterQuery(paginateQuery))
	if err != nil {
		return NewBaseApiResponse[*common.PaginateResult[entities.OrderEntity]](
			err,
			err.Error(),
			22,
			nil,
			http.StatusBadRequest,
		)
	}

	return NewBaseApiResponse[*common.PaginateResult[entities.OrderEntity]](
		nil,
		"success",
		200,
		list,
		http.StatusOK,
	)
}
