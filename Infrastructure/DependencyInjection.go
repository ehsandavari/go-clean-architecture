package Infrastructure

import (
	"GolangCodeBase/Infrastructure/Config"
	"GolangCodeBase/Infrastructure/Postgres"
	"GolangCodeBase/Infrastructure/Redis"
	"go.uber.org/fx"
)

var Module = fx.Module("Infrastructure",
	fx.Provide(
		Config.NewConfig,
		Postgres.NewPostgres,
		Redis.NewRedis,
	),
)
