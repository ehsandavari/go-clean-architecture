package Api

import (
	"GolangCodeBase/Presentation"
	"GolangCodeBase/Presentation/Common"
	"github.com/kataras/iris/v12"
	"go.uber.org/fx"
)

type SApplication struct {
}

func init() {
	Presentation.Modules = append(Presentation.Modules, fx.Provide(NewApplication))
}

func NewApplication() *SApplication {
	return &SApplication{}
}

func (r *SApplication) SetupAPI() {
	irisApp := iris.Default()
	irisApp.Use(Common.ErrorHandlerMiddleware)
	api := irisApp.Party("/api")
	r.registerOrderAPI(api)
	irisApp.Logger().Fatal(irisApp.Listen("127.0.0.1:9090"))
}
