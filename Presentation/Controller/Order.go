package Controller

import (
	"GolangCodeBase/Application/Handlers/Order/Commands/PublishOrder"
	"GolangCodeBase/Presentation/Common"
	"GolangCodeBase/Presentation/Controller/Dto"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/mehdihadeli/go-mediatr"
	"net/http"
)

// Order todo: use Validator for request body
func Order(ctx iris.Context) {
	params := &Dto.CreateOrderRequest{}
	Common.ReadJson(ctx, &params)
	_, err := mediatr.Send[PublishOrder.SPublishOrderCommand, string](ctx.Request().Context(), PublishOrder.NewSPublishOrderCommand(params.Price, params.Title))
	if err != nil {
		fmt.Println(err)
	}
	ctx.StatusCode(http.StatusOK)
}
