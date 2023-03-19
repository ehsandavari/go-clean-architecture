package v1

import (
	"github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/golang-clean-architecture/presentation/http/api/v1/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Golang Clean Architecture
//	@version		1.0
//	@description	Example API

//	@contact.name	Ehsan Davari
//	@contact.url	https://github.com/ehsandavari
//	@contact.email	ehsandavari.ir@gmail.com

//	@BasePath	/api/v1

func Setup(routerGroup *gin.RouterGroup, logger interfaces.ILogger) {
	apiRouterGroup := routerGroup.Group("/v1")
	{
		apiRouterGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v1")))
		controllers.NewOrderController(apiRouterGroup, logger)
	}
}
