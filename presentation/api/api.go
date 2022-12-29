package api

import (
	"github.com/kataras/iris/v12"
	"go.uber.org/fx"
	"golangCodeBase/presentation"
	"golangCodeBase/presentation/common"
)

func init() {
	presentation.Modules = append(presentation.Modules, fx.Provide(NewApplication))
}

type SApplication struct {
}

func NewApplication() *SApplication {
	return &SApplication{}
}

func (r *SApplication) SetupAPI() {
	irisApp := iris.Default()
	irisApp.Use(common.ErrorHandlerMiddleware)
	api := irisApp.Party("/api")
	r.registerOrderAPI(api)
	irisApp.Logger().Fatal(irisApp.Listen("127.0.0.1:9090"))
}
