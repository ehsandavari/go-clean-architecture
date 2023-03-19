package api

import (
	"github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	_ "github.com/ehsandavari/golang-clean-architecture/presentation/http/api/docs"
	v1 "github.com/ehsandavari/golang-clean-architecture/presentation/http/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

//go:generate swag fmt

func Setup(logger interfaces.ILogger) {
	engine := gin.Default()
	monitoringRouterGroup := engine.Group("/-")
	{
		monitoringRouterGroup.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })
		monitoringRouterGroup.GET("/liveness", func(c *gin.Context) { c.Status(http.StatusOK) })
		monitoringRouterGroup.GET("/readyness", func(c *gin.Context) { c.Status(http.StatusOK) })
	}
	v1.Setup(engine, logger)
	if err := engine.Run(); err != nil {
		logger.Fatal(err)
	}
}
