package Infrastructure

import (
	"GolangCodeBase/Infrastructure/Config"
	"GolangCodeBase/Infrastructure/Logger"
	"GolangCodeBase/Infrastructure/Postgres"
	"GolangCodeBase/Infrastructure/Redis"
	"go.uber.org/fx"
)

var Module = fx.Module("Infrastructure",
	fx.Provide(
		Config.NewConfig,
		func(config *Config.SConfig) (Logger.SConfig, Postgres.SConfig, Redis.SConfig) {
			return config.Service.Logger, config.Postgres, config.Redis
		},
		Logger.NewLogger,
		Postgres.NewPostgres,
		Redis.NewRedis,
	),
)
