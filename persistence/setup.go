package persistence

import (
	"github.com/ehsandavari/go-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/go-clean-architecture/infrastructure/postgres"
	"github.com/ehsandavari/go-logger"
)

type Persistence struct {
	UnitOfWork interfaces.IUnitOfWork
}

func NewPersistence(logger logger.ILogger, postgres *postgres.SPostgres) *Persistence {
	databaseContext := newDatabaseContext(logger, postgres)
	return &Persistence{
		UnitOfWork: newUnitOfWork(databaseContext),
	}
}
