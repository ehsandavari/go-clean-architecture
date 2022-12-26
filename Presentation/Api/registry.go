package Api

import (
	"GolangCodeBase/Presentation/Controller"
	"github.com/kataras/iris/v12/core/router"
)

func (r SApplication) registerOrderAPI(router router.Party) {
	router.Post("/order", Controller.Order)
}
