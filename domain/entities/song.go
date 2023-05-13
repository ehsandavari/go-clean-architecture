package entities

import (
	"github.com/ehsandavari/go-clean-architecture/domain/enums"
	"github.com/google/uuid"
)

type Song struct {
	Id              uuid.UUID
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

	Tariffs []*SongTariff
} //@name Song

func NewSong(id uuid.UUID, userId uint, title string, genre enums.SongGenre, sense enums.SongSense, subject enums.SongSubject, beat byte, language enums.SongLanguage, minPrice uint, maxPrice uint, price uint, displayableText string, fullText string, status enums.SongStatus, base Base) *Song {
	return &Song{
		Id:              id,
		UserId:          userId,
		Title:           title,
		Genre:           genre,
		Sense:           sense,
		Subject:         subject,
		Beat:            beat,
		Language:        language,
		MinPrice:        minPrice,
		MaxPrice:        maxPrice,
		Price:           price,
		DisplayableText: displayableText,
		FullText:        fullText,
		Status:          status,
		Base:            base,
	}
}

func (r *Song) AddTariff(tariff *SongTariff) {
	r.Tariffs = append(r.Tariffs, tariff)
}
