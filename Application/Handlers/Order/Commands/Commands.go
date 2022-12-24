package Commands

import (
	ApplicationInterfaces "GolangCodeBase/Application/Common/Interfaces"
	DomainInterfaces "GolangCodeBase/Domain/Interfaces"
	"GolangCodeBase/Infrastructure/Config"
)

type SOrderHandlerCommands struct {
	sConfig     *Config.SConfig
	iLogger     ApplicationInterfaces.ILogger
	iUnitOfWork DomainInterfaces.IUnitOfWork
	iRedis      ApplicationInterfaces.IRedis
}

func NewOrderHandlerCommands(sOrderHandlerCommandsParams sOrderHandlerCommandsParams) ApplicationInterfaces.IOrderHandlerCommands {
	return &SOrderHandlerCommands{
		sConfig:     sOrderHandlerCommandsParams.SConfig,
		iLogger:     sOrderHandlerCommandsParams.ILogger,
		iUnitOfWork: sOrderHandlerCommandsParams.IUnitOfWork,
		iRedis:      sOrderHandlerCommandsParams.IRedis,
	}
}
