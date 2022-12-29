package api

import (
	"github.com/kataras/iris/v12/core/router"
	"golangCodeBase/presentation/controller"
)

func (r SApplication) registerOrderAPI(router router.Party) {
	router.Post("/order", controller.Order)
}
