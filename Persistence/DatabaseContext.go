package Persistence

import (
	"GolangCodeBase/Infrastructure/Postgres"
	"go.uber.org/fx"
)

type sDatabaseContext struct {
	postgres *Postgres.SPostgres
}

func init() {
	Modules = append(Modules, fx.Provide(NewDatabaseContext))
}

func NewDatabaseContext(postgres *Postgres.SPostgres) (*sDatabaseContext, error) {
	return &sDatabaseContext{
		postgres: postgres,
	}, nil
}
