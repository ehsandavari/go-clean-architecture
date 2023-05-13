package createSong

import (
	"context"
	"github.com/ehsandavari/go-clean-architecture/application"
	"github.com/ehsandavari/go-clean-architecture/application/common"
	"github.com/ehsandavari/go-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/go-clean-architecture/domain/entities"
	"github.com/ehsandavari/go-logger"
	"github.com/ehsandavari/go-mediator"
	"github.com/ehsandavari/go-tracer"
	"github.com/google/uuid"
	"log"
	"time"
)

func init() {
	application.Handlers = append(application.Handlers, func(application *application.Application) {
		if err := mediator.RegisterRequestHandler[SCreateSongCommand, *entities.Song](newCreateSongCommandHandler(application.ILogger, application.ITracer, application.IUnitOfWork)); err != nil {
			log.Fatalln(err)
		}
	})
}

type SCreateSongCommandHandler struct {
	iLogger     logger.ILogger
	iTracer     tracer.ITracer
	iUnitOfWork interfaces.IUnitOfWork
}

func newCreateSongCommandHandler(
	iLogger logger.ILogger,
	iTracer tracer.ITracer,
	iUnitOfWork interfaces.IUnitOfWork,
) SCreateSongCommandHandler {
	return SCreateSongCommandHandler{
		iLogger:     iLogger,
		iTracer:     iTracer,
		iUnitOfWork: iUnitOfWork,
	}
}

func (r SCreateSongCommandHandler) Handle(ctx context.Context, command SCreateSongCommand) (*entities.Song, mediator.IError) {
	spanTracer := r.iTracer.Tracer("test")
	spanTracer.Start(ctx, "testHandle")
	defer spanTracer.End()

	timeNow := time.Now()
	songEntity := entities.NewSong(
		uuid.New(),
		command.UserId,
		command.Title,
		command.Genre,
		command.Sense,
		command.Subject,
		command.Beat,
		command.Language,
		command.Price,
		command.MaxPrice,
		command.Price,
		command.DisplayableText,
		command.FullText,
		command.Status,
		entities.NewBase(timeNow, timeNow),
	)
	for _, tariffId := range command.TariffIds {
		//tariff, err := r.iUnitOfWork.SongTariffRepository().FirstById(ctx, tariffId)
		//if err != nil {
		//	return nil, common.ErrorInternal
		//}
		songEntity.AddTariff(entities.NewSongTariff(songEntity.Id, tariffId, entities.NewBase(timeNow, timeNow)))
	}
	createdSongNumber, err := r.iUnitOfWork.SongRepository().Create(ctx, songEntity)
	if err != nil {
		return nil, common.ErrorInternal
	}
	if createdSongNumber == 0 {
		r.iLogger.Warn("song not created")
		return nil, common.ErrorInternal
	}
	return songEntity, nil
}
