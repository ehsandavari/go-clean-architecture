package persistence

import (
	"go.uber.org/fx"
	"golangCodeBase/infrastructure/postgres"
)

func init() {
	Modules = append(Modules, fx.Provide(NewDatabaseContext))
}

type sDatabaseContext struct {
	postgres *postgres.SPostgres
}

func NewDatabaseContext(postgres *postgres.SPostgres) (*sDatabaseContext, error) {
	return &sDatabaseContext{
		postgres: postgres,
	}, nil
}
