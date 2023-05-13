package grpc

import (
	"github.com/ehsandavari/go-clean-architecture/infrastructure/config"
	bankAccountService "github.com/ehsandavari/go-clean-architecture/presentation/grpc/proto/bankAccount"
	"github.com/ehsandavari/go-logger"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type SGrpc struct {
	server  *grpc.Server
	sConfig *config.SConfig
	iLogger logger.ILogger
}

func NewSGrpc(sConfig *config.SConfig, iLogger logger.ILogger) *SGrpc {
	var sGrpc SGrpc
	sGrpc.sConfig = sConfig
	if sConfig.Service.Grpc.IsEnabled {
		sGrpc.server = grpc.NewServer()
		sGrpc.iLogger = iLogger
	}
	return &sGrpc
}

func (r *SGrpc) Start() {
	if r.sConfig.Service.Grpc.IsEnabled {
		netListener, err := net.Listen("tcp", ":"+r.sConfig.Service.Grpc.Port)
		if err != nil {
			log.Fatal("error in net listen ", err)
		}
		bankAccountService.RegisterBankAccountServiceServer(r.server, NewGrpcService())
		grpc_prometheus.Register(r.server)

		if r.sConfig.Service.Grpc.IsDevelopment {
			reflection.Register(r.server)
		}

		go func() {
			err := r.server.Serve(netListener)
			if err != nil {
				r.iLogger.Fatal("error in serve grpc server", r.server.Serve(netListener))
			}
		}()
		r.iLogger.Infof("grpc server start on port: %s, app info: %+v", r.sConfig.Service.Grpc.Port, r.server.GetServiceInfo())
	}
}

func (r *SGrpc) Stop() {
	if r.sConfig.Service.Grpc.IsEnabled {
		r.server.GracefulStop()
	}
}
