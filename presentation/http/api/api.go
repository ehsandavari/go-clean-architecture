package api

import (
	"github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	_ "github.com/ehsandavari/golang-clean-architecture/presentation/http/api/docs"
	"github.com/ehsandavari/golang-clean-architecture/presentation/http/api/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func Setup(logger interfaces.ILogger) {
	engine := gin.Default()
	engine.Use(cors.Default())
	monitoringRouterGroup := engine.Group("/-")
	{
		monitoringRouterGroup.GET("/health", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })
		monitoringRouterGroup.GET("/liveness", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })
		monitoringRouterGroup.GET("/readyness", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })
		monitoringRouterGroup.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}

	apiRouterGroup := engine.Group("/api")
	{
		v1.Setup(apiRouterGroup, logger)
	}
	if err := engine.Run(); err != nil {
		logger.Fatal(err)
	}
}
