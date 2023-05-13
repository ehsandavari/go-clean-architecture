package application

import (
	"github.com/ehsandavari/go-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/go-clean-architecture/infrastructure/config"
	"github.com/ehsandavari/go-logger"
	"github.com/ehsandavari/go-tracer"
)

type Application struct {
	SConfig     *config.SConfig
	ILogger     logger.ILogger
	ITracer     tracer.ITracer
	IUnitOfWork interfaces.IUnitOfWork
}

func NewApplication(sConfig *config.SConfig, iLogger logger.ILogger, iTracer tracer.ITracer, iUnitOfWork interfaces.IUnitOfWork) *Application {
	return &Application{
		SConfig:     sConfig,
		ILogger:     iLogger,
		ITracer:     iTracer,
		IUnitOfWork: iUnitOfWork,
	}
}

var Handlers []func(application *Application)

func (r *Application) Setup() {
	for _, handler := range Handlers {
		handler(r)
	}
}
