package api

import (
	"github.com/ehsandavari/golang-clean-architecture/presentation/controller"
	"github.com/kataras/iris/v12/core/router"
)

func (r SApplication) registerOrderAPI(router router.Party) {
	router.Post("/order", controller.Order)
}
