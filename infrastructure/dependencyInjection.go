package infrastructure

import (
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/config"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/logger"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/postgres"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/redis"
	"go.uber.org/fx"
)

var Modules = fx.Module("infrastructure",
	fx.Provide(
		config.NewConfig,
		func(config *config.SConfig) (logger.SConfig, postgres.SConfig, redis.SConfig) {
			return config.Service.Logger, config.Postgres, config.Redis
		},
		logger.NewLogger,
		postgres.NewPostgres,
		redis.NewRedis,
	),
)
