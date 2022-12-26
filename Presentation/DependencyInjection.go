package Presentation

import (
	"GolangCodeBase/Presentation/Api"
	"go.uber.org/fx"
)

var Module = fx.Module("Presentation",
	fx.Provide(
		Api.NewApplication,
	),
)
