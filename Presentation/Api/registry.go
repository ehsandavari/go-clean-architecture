package Api

import (
	"github.com/kataras/iris/v12/core/router"
)

func (r SApplication) registerOrderAPI(router router.Party) {
	router.Post("/order", r.OrderController.Order)
}
