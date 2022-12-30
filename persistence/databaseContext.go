package persistence

import (
	"go.uber.org/fx"
	"golangCodeBase/infrastructure/postgres"
)

func init() {
	Modules = append(Modules, fx.Provide(newDatabaseContext))
}

type SDatabaseContext struct {
	Postgres *postgres.SPostgres
}

func newDatabaseContext(postgres *postgres.SPostgres) (*SDatabaseContext, error) {
	return &SDatabaseContext{
		Postgres: postgres,
	}, nil
}
