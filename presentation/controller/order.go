package controller

import (
	"fmt"
	"github.com/ehsandavari/golang-clean-architecture/application/handlers/order/commands/publishOrder"
	"github.com/ehsandavari/golang-clean-architecture/presentation/common"
	"github.com/ehsandavari/golang-clean-architecture/presentation/controller/dto"
	"net/http"
)

// Order todo: use Validator for request body
func Order(ctx iris.Context) {
	params := &dto.CreateOrderRequest{}
	common.ReadJson(ctx, &params)
	_, err := mediator.Send[publishOrder.SPublishOrderCommand, string](ctx.Request().Context(), publishOrder.NewSPublishOrderCommand(params.Price, params.Title))
	if err != nil {
		fmt.Println(err)
	}
	ctx.StatusCode(http.StatusOK)
}
