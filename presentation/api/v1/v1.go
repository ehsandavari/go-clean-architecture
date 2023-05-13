package v1

import (
	"github.com/ehsandavari/go-clean-architecture/presentation/api/v1/controllers"
	"github.com/ehsandavari/go-logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			api
//	@version		1.0
//	@description	Example Api

//	@contact.name	Ehsan Davari
//	@contact.url	https://github.com/ehsandavari
//	@contact.email	ehsandavari.ir@gmail.com

//	@BasePath	/api/v1

func Setup(routerGroup *gin.RouterGroup, logger logger.ILogger) {
	apiRouterGroup := routerGroup.Group("/v1")
	{
		apiRouterGroup.GET("/swagger/*any", ginSwagger.WrapHandler(
			swaggerFiles.NewHandler(),
			ginSwagger.InstanceName("v1"),
		))
		controllers.NewSongController(apiRouterGroup, logger)
	}
}
