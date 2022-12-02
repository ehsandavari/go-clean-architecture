package Common

import (
	"github.com/kataras/iris/v12"
)

func ReadJson(context iris.Context, structure interface{}) {
	if err := context.ReadJSON(structure); err != nil {
		panic(errorResponse(err))
	}
}
