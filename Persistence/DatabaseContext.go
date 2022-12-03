package Persistence

import (
	"GolangCodeBase/Infrastructure/Postgres"
	"go.uber.org/fx"
)

type sDatabaseContext struct {
	postgres *Postgres.SPostgres
}

type sDatabaseContextParams struct {
	fx.In
	Postgres *Postgres.SPostgres
}

func NewDatabaseContext(databaseContextParams sDatabaseContextParams) (*sDatabaseContext, error) {
	return &sDatabaseContext{
		postgres: databaseContextParams.Postgres,
	}, nil
}
