package Api

import (
	"GolangCodeBase/Presentation/Common"
	"GolangCodeBase/Presentation/Controller"
	"github.com/kataras/iris/v12"
)

type SApplication struct {
	OrderController Controller.SOrderController
}

func NewApplication(sOrderController Controller.SOrderController) *SApplication {
	return &SApplication{
		OrderController: sOrderController,
	}
}

func (r *SApplication) SetupAPI() {
	irisApp := iris.Default()
	irisApp.Use(Common.ErrorHandlerMiddleware)
	api := irisApp.Party("/api")
	r.registerOrderAPI(api)
	irisApp.Logger().Fatal(irisApp.Listen("127.0.0.1:8080"))
}
