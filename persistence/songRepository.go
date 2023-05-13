package persistence

import (
	"github.com/ehsandavari/go-clean-architecture/domain/entities"
	"github.com/ehsandavari/go-clean-architecture/infrastructure/postgres/models"
)

type sSongRepository struct {
	sGenericRepository[models.Song, entities.Song]
}

func newSongRepository(db *sDatabaseContext) sSongRepository {
	return sSongRepository{
		sGenericRepository: newGenericRepository[models.Song, entities.Song](db),
	}
}
