package api

import (
	"github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/config"
	_ "github.com/ehsandavari/golang-clean-architecture/presentation/http/api/docs"
	"github.com/ehsandavari/golang-clean-architecture/presentation/http/api/middlewares"
	"github.com/ehsandavari/golang-clean-architecture/presentation/http/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
	"net/http"
)

type SApi struct {
	fx.In
	Config *config.SConfig
	Logger interfaces.ILogger
}

func Setup(api SApi) {
	gin.SetMode(api.Config.Service.Http.Mode)
	engine := gin.Default()
	engine.Use(
		middlewares.Cors(),
		middlewares.I18n(),
	)

	monitoringRouterGroup := engine.Group("/-")
	{
		monitoringRouterGroup.GET("/health", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })
		monitoringRouterGroup.GET("/liveness", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })
		monitoringRouterGroup.GET("/readyness", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })
		monitoringRouterGroup.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}

	apiRouterGroup := engine.Group("/api")
	{
		v1.Setup(apiRouterGroup, api.Logger)
	}

	go func() {
		if err := engine.Run(api.Config.Service.Http.Host + ":" + api.Config.Service.Http.Port); err != nil {
			api.Logger.Fatal(err)
		}
	}()
}
