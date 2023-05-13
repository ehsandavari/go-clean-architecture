package persistence

import (
	"github.com/ehsandavari/go-clean-architecture/infrastructure/postgres"
	"github.com/ehsandavari/go-logger"
)

type sDatabaseContext struct {
	Logger   logger.ILogger
	Postgres *postgres.SPostgres
}

func newDatabaseContext(logger logger.ILogger, postgres *postgres.SPostgres) *sDatabaseContext {
	return &sDatabaseContext{
		Logger:   logger,
		Postgres: postgres,
	}
}
