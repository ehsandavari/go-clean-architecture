package persistence

import (
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/postgres"
)

type SDatabaseContext struct {
	Postgres *postgres.SPostgres
}

func newDatabaseContext(postgres *postgres.SPostgres) (*SDatabaseContext, error) {
	return &SDatabaseContext{
		Postgres: postgres,
	}, nil
}
