package controllers

import (
	"errors"
	"github.com/ehsandavari/go-mediator"
	"github.com/ehsandavari/golang-clean-architecture/application/handlers/order/commands/publishOrder"
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
//	@Success		200	{object}	BaseApiResponse
//	@Router			/ [get]
func (c *Controller) GetMap(ctx *gin.Context) BaseApiResponse {
	_, err := mediator.Send[publishOrder.SPublishOrderCommand, string](ctx, publishOrder.NewSPublishOrderCommand(111, "params.Title"))
	if err != nil {
		return NewBaseApiResponse(
			err,
			err.Error(),
			22,
			http.StatusBadRequest,
		)
	}

	return NewBaseApiResponse(
		errors.New("asdsadasd"),
		"asdasdasd",
		11,
		http.StatusOK,
	)
}
