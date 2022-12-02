package Application

import (
	"GolangCodeBase/Application/Handlers/Order/Commands"
	"go.uber.org/fx"
)

var Module = fx.Module("Application",
	fx.Provide(
		Commands.NewOrderHandlerCommands,
	),
)
