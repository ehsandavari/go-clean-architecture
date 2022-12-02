package Repositories

import (
	"GolangCodeBase/Infrastructure/Postgres"
	"go.uber.org/fx"
)

type SDatabaseContext struct {
	Postgres *Postgres.SPostgres
}

type sDatabaseContextParams struct {
	fx.In
	Postgres *Postgres.SPostgres
}

func NewDatabaseContext(databaseContextParams sDatabaseContextParams) (*SDatabaseContext, error) {
	return &SDatabaseContext{
		Postgres: databaseContextParams.Postgres,
	}, nil
}
