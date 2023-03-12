package v1

import (
	"github.com/ehsandavari/golang-clean-architecture/application/common"
	"github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/golang-clean-architecture/domain/entities"
	"github.com/ehsandavari/golang-clean-architecture/presentation/http/api/v1/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//go:generate swag i -g setup.go -o ../docs -ot go --instanceName v1 --parseDependency

//	@title			Golang Clean Architecture
//	@version		1.0
//	@description	Example API

//	@contact.name	Ehsan Davari
//	@contact.url	https://github.com/ehsandavari
//	@contact.email	ehsandavari.ir@gmail.com

//	@BasePath	/api/v1

func Setup(engine *gin.Engine, logger interfaces.ILogger) {
	engine.Use(cors.Default())
	setupSwagger(engine)
	controller := controllers.NewController(logger)
	engine.POST("/api/v1", controllers.AppController[*common.PaginateResult[entities.OrderEntity]](controller.GetMap).Handler())
}

func setupSwagger(engine *gin.Engine) {
	engine.GET("/swagger/v1/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v1")))
}
