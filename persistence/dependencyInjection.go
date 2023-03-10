package persistence

import (
	"go.uber.org/fx"
)

var Modules = fx.Module("persistence",
	fx.Provide(
		newDatabaseContext,
		newUnitOfWork,
	),
)
