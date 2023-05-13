package infrastructure

import (
	"github.com/ehsandavari/go-clean-architecture/infrastructure/config"
	"github.com/ehsandavari/go-clean-architecture/infrastructure/postgres"
	"github.com/ehsandavari/go-logger"
	"github.com/ehsandavari/go-tracer"
)

type Infrastructure struct {
	SConfig   *config.SConfig
	ILogger   logger.ILogger
	SPostgres *postgres.SPostgres
	ITracer   tracer.ITracer
}

func NewInfrastructure() *Infrastructure {
	sConfig := config.NewConfig()
	return &Infrastructure{
		SConfig:   sConfig,
		ILogger:   logger.NewLogger(sConfig.Logger.Level, sConfig.Logger.Mode, sConfig.Logger.Encoder),
		SPostgres: postgres.NewPostgres(sConfig.Postgres),
		ITracer: tracer.NewTracer(
			sConfig.Tracer.IsEnabled,
			sConfig.Tracer.Host,
			sConfig.Tracer.Port,
			sConfig.Service.Id,
			sConfig.Service.Name,
			sConfig.Service.Version,
			sConfig.Service.Mode.String(),
			sConfig.Tracer.Sampler,
			sConfig.Tracer.UseStdout,
		),
	}
}

func (r *Infrastructure) Close() {
	if err := r.ILogger.Sync(); err != nil {
		r.ILogger.Error("error in sync logger : ", err)
	}

	if err := r.SPostgres.Close(); err != nil {
		r.ILogger.Error("error in close postgres : ", err)
	}

	if err := r.ITracer.Shutdown(); err != nil {
		r.ILogger.Error("error in sync logger : ", err)
	}
}
