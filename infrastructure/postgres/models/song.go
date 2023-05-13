package models

import (
	"github.com/ehsandavari/go-clean-architecture/domain/entities"
	"github.com/ehsandavari/go-clean-architecture/domain/enums"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Song struct {
	Id              uuid.UUID `gorm:"type:uuid;"`
	UserId          uint
	Title           string
	Genre           enums.SongGenre
	Sense           enums.SongSense
	Subject         enums.SongSubject
	Beat            byte
	Language        enums.SongLanguage
	MinPrice        uint
	MaxPrice        uint
	Price           uint
	DisplayableText string
	FullText        string
	Status          enums.SongStatus
	Base

	Tariffs []SongTariff
}

func (r *Song) BeforeCreate(*gorm.DB) error {
	r.Id = uuid.New()
	return nil
}

func (r Song) ToEntity() *entities.Song {
	song := entities.NewSong(
		r.Id,
		r.UserId,
		r.Title,
		r.Genre,
		r.Sense,
		r.Subject,
		r.Beat,
		r.Language,
		r.MinPrice,
		r.MaxPrice,
		r.Price,
		r.DisplayableText,
		r.FullText,
		r.Status,
		r.Base.ToEntity(),
	)
	for _, tariff := range r.Tariffs {
		song.AddTariff(tariff.ToEntity())
	}
	return song
}

func (r Song) FromEntity(entity *entities.Song) any {
	r.Id = entity.Id
	r.UserId = entity.UserId
	r.Title = entity.Title
	r.Genre = entity.Genre
	r.Sense = entity.Sense
	r.Subject = entity.Subject
	r.Beat = entity.Beat
	r.Language = entity.Language
	r.MinPrice = entity.MinPrice
	r.MaxPrice = entity.MaxPrice
	r.Price = entity.Price
	r.DisplayableText = entity.DisplayableText
	r.FullText = entity.FullText
	r.Status = entity.Status
	r.Base.FromEntity(entity.Base)
	for _, tariff := range entity.Tariffs {
		r.Tariffs = append(r.Tariffs, new(SongTariff).FromEntity(tariff).(SongTariff))
	}
	return r
}
