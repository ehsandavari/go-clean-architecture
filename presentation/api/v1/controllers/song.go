package controllers

import (
	"github.com/ehsandavari/go-clean-architecture/application/handlers/order/commands/createSong"
	"github.com/ehsandavari/go-clean-architecture/domain/entities"
	"github.com/ehsandavari/go-clean-architecture/domain/enums"
	"github.com/ehsandavari/go-clean-architecture/presentation/api/v1/dtos"
	"github.com/ehsandavari/go-logger"
	"github.com/ehsandavari/go-mediator"
	"github.com/gin-gonic/gin"
)

type sSongController struct {
	sBaseController
}

func NewSongController(routerGroup *gin.RouterGroup, logger logger.ILogger) {
	songController := sSongController{
		sBaseController: newBaseController(logger),
	}

	routerGroup = routerGroup.Group("/song")
	{
		routerGroup.POST("/create", baseController[*dtos.CreateSongRequest, *entities.Song](songController.create).Handle(songController.iLogger))
	}
}

// GetSongs
//
//	@Tags		song
//	@Summary	create song
//	@Accept		json
//	@Produce	json
//	@Param		Accept-Language	header		string				false	"some description"	Enums(en, fa)
//	@Param		params			body		CreateSongRequest	false	"Query Params"
//	@Success	200				{object}	BaseApiResponse[entities.Song]
//	@Failure	400				{object}	BaseApiResponse[ApiError]
//	@Failure	500				{object}	BaseApiResponse[FilterQuery]
//	@Router		/song/create [POST]
func (r *sSongController) create(ctx *gin.Context, dto *dtos.CreateSongRequest) (*entities.Song, error) {
	entity, err := mediator.Send[createSong.SCreateSongCommand, *entities.Song](ctx, createSong.NewSCreateSongCommand(
		1,
		dto.Title,
		dto.Genre,
		dto.Sense,
		dto.Subject,
		dto.Beat,
		dto.Language,
		dto.Price,
		dto.MaxPrice,
		dto.Price,
		dto.DisplayableText,
		dto.FullText,
		enums.SongStatusPending,
		dto.Tariffs,
	))
	if err != nil {
		return nil, err
	}

	return entity, nil
}
