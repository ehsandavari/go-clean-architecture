package Persistence

import (
	"GolangCodeBase/Infrastructure/Postgres"
	"go.uber.org/fx"
)

func init() {
	Modules = append(Modules, fx.Provide(NewDatabaseContext))
}

type sDatabaseContext struct {
	postgres *Postgres.SPostgres
}

func NewDatabaseContext(postgres *Postgres.SPostgres) (*sDatabaseContext, error) {
	return &sDatabaseContext{
		postgres: postgres,
	}, nil
}
