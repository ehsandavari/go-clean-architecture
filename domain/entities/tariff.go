package entities

import (
	"github.com/google/uuid"
)

type SongTariff struct {
	SongId   uuid.UUID
	TariffId uuid.UUID
	Base
} //@name SongTariff

func NewSongTariff(songId uuid.UUID, tariffId uuid.UUID, base Base) *SongTariff {
	return &SongTariff{
		SongId:   songId,
		TariffId: tariffId,
		Base:     base,
	}
}
