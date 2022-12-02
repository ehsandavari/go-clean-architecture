package Persistence

import (
	"GolangCodeBase/Persistence/Repositories"
	"go.uber.org/fx"
)

var Module = fx.Module("Persistence",
	fx.Provide(
		Repositories.NewDatabaseContext,
		Repositories.NewUnitOfWork,
	),
)
