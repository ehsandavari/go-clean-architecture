package api

import (
	"context"
	"github.com/ehsandavari/go-clean-architecture/infrastructure/config"
	_ "github.com/ehsandavari/go-clean-architecture/presentation/api/docs"
	"github.com/ehsandavari/go-clean-architecture/presentation/api/middlewares"
	"github.com/ehsandavari/go-clean-architecture/presentation/api/v1"
	"github.com/ehsandavari/go-logger"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type SApi struct {
	server  *http.Server
	sConfig *config.SConfig
	iLogger logger.ILogger
}

func NewSApi(sConfig *config.SConfig, iLogger logger.ILogger) *SApi {
	var sApi SApi
	sApi.sConfig = sConfig
	if sConfig.Service.Api.IsEnabled {
		sApi.server = &http.Server{
			Addr: sConfig.Service.Api.Host + ":" + sConfig.Service.Api.Port,
		}
		sApi.iLogger = iLogger
	}
	return &sApi
}

func (r *SApi) Start() {
	if r.sConfig.Service.Api.IsEnabled {
		gin.SetMode(r.sConfig.Service.Api.Mode)
		engine := gin.Default()
		engine.Use(
			middlewares.Cors(),
			middlewares.I18n(),
			middlewares.RequestId(),
		)

		monitoringRouterGroup := engine.Group("/-")
		{
			monitoringRouterGroup.GET("/health", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })
			monitoringRouterGroup.GET("/liveness", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })
			monitoringRouterGroup.GET("/readiness", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })
			monitoringRouterGroup.GET("/metrics", gin.WrapH(promhttp.Handler()))
		}

		apiRouterGroup := engine.Group("/api")
		{
			v1.Setup(apiRouterGroup, r.iLogger)
		}

		go func() {
			r.server.Handler = engine.Handler()
			if err := r.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				r.iLogger.Fatal("error in serve api server", err)
			}
		}()
		r.iLogger.Info("api server start on port : ", r.sConfig.Service.Api.Port)
	}
}

func (r *SApi) Stop() {
	if r.sConfig.Service.Api.IsEnabled {
		err := r.server.Shutdown(context.Background())
		if err != nil {
			r.iLogger.Error("error in shutdown api server", err)
		}
	}
}
