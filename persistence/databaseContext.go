package persistence

import (
	"github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/postgres"
)

type SDatabaseContext struct {
	Logger   interfaces.ILogger
	Postgres *postgres.SPostgres
}

func newDatabaseContext(postgres *postgres.SPostgres, logger interfaces.ILogger) (*SDatabaseContext, error) {
	return &SDatabaseContext{
		Logger:   logger,
		Postgres: postgres,
	}, nil
}
