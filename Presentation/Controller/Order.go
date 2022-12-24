package Controller

import (
	ApplicationInterfaces "GolangCodeBase/Application/Common/Interfaces"
	DomainEntities "GolangCodeBase/Domain/Entities"
	"GolangCodeBase/Presentation/Common"
	"GolangCodeBase/Presentation/Controller/Dto"
	"github.com/kataras/iris/v12"
	"net/http"
)

type SOrderController struct {
	iOrderCommand ApplicationInterfaces.IOrderHandlerCommands
}

func NewOrderController(iOrderCommand ApplicationInterfaces.IOrderHandlerCommands) SOrderController {
	return SOrderController{
		iOrderCommand: iOrderCommand,
	}
}

// Order todo: use Validator for request body
func (sOrderController SOrderController) Order(ctx iris.Context) {
	params := &Dto.CreateOrderRequest{}
	Common.ReadJson(ctx, &params)
	sOrderController.iOrderCommand.PublishOrderCommand(ctx.Request().Context(), DomainEntities.OrderEntity{
		Price: 0,
		Title: "",
	})
	ctx.StatusCode(http.StatusOK)
}
