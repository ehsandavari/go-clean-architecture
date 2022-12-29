package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/mehdihadeli/go-mediatr"
	"golangCodeBase/application/handlers/order/commands/publishOrder"
	"golangCodeBase/presentation/common"
	"golangCodeBase/presentation/controller/dto"
	"net/http"
)

// Order todo: use Validator for request body
func Order(ctx iris.Context) {
	params := &dto.CreateOrderRequest{}
	common.ReadJson(ctx, &params)
	_, err := mediatr.Send[publishOrder.SPublishOrderCommand, string](ctx.Request().Context(), publishOrder.NewSPublishOrderCommand(params.Price, params.Title))
	if err != nil {
		fmt.Println(err)
	}
	ctx.StatusCode(http.StatusOK)
}
