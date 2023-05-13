package persistence

import (
	"github.com/ehsandavari/go-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/go-clean-architecture/infrastructure/postgres"
)

type sUnitOfWork struct {
	databaseContext *sDatabaseContext
	songRepository  interfaces.ISongRepository
}

func newUnitOfWork(databaseContext *sDatabaseContext) interfaces.IUnitOfWork {
	return &sUnitOfWork{
		databaseContext: databaseContext,
		songRepository:  newSongRepository(databaseContext),
	}
}

func (r *sUnitOfWork) SongRepository() interfaces.ISongRepository {
	return r.songRepository
}

func (r *sUnitOfWork) Do(unitOfWorkBlock func(interfaces.IUnitOfWork) error) error {
	return r.databaseContext.Postgres.Transaction(func(transaction *postgres.SPostgres) error {
		r.databaseContext.Postgres = transaction
		return unitOfWorkBlock(r)
	})
}
