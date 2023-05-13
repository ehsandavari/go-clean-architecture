package presentation

import (
	"github.com/ehsandavari/go-clean-architecture/infrastructure/config"
	"github.com/ehsandavari/go-clean-architecture/presentation/api"
	"github.com/ehsandavari/go-clean-architecture/presentation/grpc"
	"github.com/ehsandavari/go-logger"
)

type Presentation struct {
	sApi  *api.SApi
	sGrpc *grpc.SGrpc
}

func NewPresentation(sConfig *config.SConfig, iLogger logger.ILogger) *Presentation {
	return &Presentation{
		sApi:  api.NewSApi(sConfig, iLogger),
		sGrpc: grpc.NewSGrpc(sConfig, iLogger),
	}
}

func (r *Presentation) Setup() {
	r.sApi.Start()
	r.sGrpc.Start()
}

func (r *Presentation) Close() {
	r.sApi.Stop()
	r.sGrpc.Stop()
}
