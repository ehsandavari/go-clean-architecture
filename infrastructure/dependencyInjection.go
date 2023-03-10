package infrastructure

import (
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/config"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/logger"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/postgres"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/redis"
	"go.uber.org/fx"
)

var Modules = fx.Module("infrastructure",
	fx.Invoke(
		config.NewConfig,
		logger.NewLogger,
		postgres.NewPostgres,
		redis.NewRedis,
	),
)
