package Presentation

import (
	"GolangCodeBase/Presentation/Api"
	"GolangCodeBase/Presentation/Controller"
	"go.uber.org/fx"
)

var Module = fx.Module("Presentation",
	fx.Provide(
		Controller.NewOrderController,
		Api.NewApplication,
	),
)
