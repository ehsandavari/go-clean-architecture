package dtos

import (
	"github.com/ehsandavari/go-clean-architecture/domain/enums"
	"github.com/google/uuid"
)

type CreateSongRequest struct {
	Title           string             `binding:"required"`
	Genre           enums.SongGenre    `binding:"required"`
	Sense           enums.SongSense    `binding:"required"`
	Subject         enums.SongSubject  `binding:"required"`
	Beat            byte               `binding:"required,min=10"`
	Language        enums.SongLanguage `binding:"required"`
	MinPrice        uint               `binding:"required"`
	MaxPrice        uint               `binding:"required"`
	Price           uint               `binding:"required"`
	DisplayableText string             `binding:"required"`
	FullText        string             `binding:"required"`
	Tariffs         []uuid.UUID        `binding:"required"`
} //@name CreateSongRequest
