package Persistence

import (
	"go.uber.org/fx"
)

var Module = fx.Module("Persistence",
	fx.Provide(
		NewDatabaseContext,
		NewUnitOfWork,
	),
)
