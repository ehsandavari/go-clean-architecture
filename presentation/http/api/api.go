package api

import (
	"github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	_ "github.com/ehsandavari/golang-clean-architecture/presentation/http/api/docs"
	v1 "github.com/ehsandavari/golang-clean-architecture/presentation/http/api/v1"
	v2 "github.com/ehsandavari/golang-clean-architecture/presentation/http/api/v2"
	"github.com/gin-gonic/gin"
)

//go:generate swag fmt

func Setup(logger interfaces.ILogger) {
	engine := gin.Default()
	v1.Setup(engine, logger)
	v2.Setup(engine)
	if err := engine.Run(); err != nil {
		logger.Fatal(err)
	}
}
