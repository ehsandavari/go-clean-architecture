package infrastructure

import (
	"github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/config"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/logger"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/postgres"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/redis"
	"go.uber.org/fx"
)

var Modules = fx.Module("infrastructure",
	fx.Provide(
		config.NewConfig,
		func(lc fx.Lifecycle, config *config.SConfig) (interfaces.ILogger, *postgres.SPostgres, interfaces.IRedis) {
			return logger.NewLogger(config.Service.Logger),
				postgres.NewPostgres(lc, config.Postgres),
				redis.NewRedis(lc, config.Redis)
		},
	),
)
