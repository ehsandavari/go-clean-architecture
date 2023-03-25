package presentation

import (
	"github.com/ehsandavari/golang-clean-architecture/presentation/http/api"
	"go.uber.org/fx"
)

var Modules = fx.Module("presentation",
	fx.Invoke(api.Setup),
)
