package models

import (
	"github.com/ehsandavari/go-clean-architecture/domain/entities"
	"github.com/google/uuid"
)

type SongTariff struct {
	SongId   uuid.UUID `gorm:"type:uuid;"`
	TariffId uuid.UUID `gorm:"type:uuid;"`
	Base
}

func (r SongTariff) ToEntity() *entities.SongTariff {
	return entities.NewSongTariff(
		r.SongId,
		r.TariffId,
		r.Base.ToEntity(),
	)
}

func (r SongTariff) FromEntity(entity *entities.SongTariff) any {
	r.SongId = entity.SongId
	r.TariffId = entity.TariffId
	r.Base.FromEntity(entity.Base)
	return r
}
