package persistence

import (
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/postgres"
	"go.uber.org/fx"
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
