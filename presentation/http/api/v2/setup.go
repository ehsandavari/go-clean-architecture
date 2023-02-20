package v2

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//go:generate swag i -g setup.go -o ../docs -ot go --instanceName v2

//	@title			Golang Clean Architecture
//	@version		2.0
//	@description	Example API

//	@contact.name	Ehsan Davari
//	@contact.url	https://github.com/ehsandavari
//	@contact.email	ehsandavari.ir@gmail.com

//	@BasePath	/api/v2

func Setup(engine *gin.Engine) {
	engine.Use(cors.Default())
	setupSwagger(engine)
}

func setupSwagger(engine *gin.Engine) {
	engine.GET("/swagger/v2/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v2")))
}
