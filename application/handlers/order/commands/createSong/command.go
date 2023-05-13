package createSong

import (
	"github.com/ehsandavari/go-clean-architecture/domain/enums"
	"github.com/google/uuid"
)

type SCreateSongCommand struct {
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
	TariffIds       []uuid.UUID
}

func NewSCreateSongCommand(userId uint, title string, genre enums.SongGenre, sense enums.SongSense, subject enums.SongSubject, beat byte, language enums.SongLanguage, minPrice uint, maxPrice uint, price uint, displayableText string, fullText string, status enums.SongStatus, tariffIds []uuid.UUID) SCreateSongCommand {
	return SCreateSongCommand{
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
		TariffIds:       tariffIds,
	}
}
