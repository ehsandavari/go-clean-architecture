package Commands

import (
	ApplicationInterfaces "GolangCodeBase/Application/Common/Interfaces"
	DomainInterfaces "GolangCodeBase/Domain/Interfaces"
	"GolangCodeBase/Infrastructure/Config"
	"go.uber.org/fx"
)

type SOrderHandlerCommands struct {
	sConfig     *Config.SConfig
	iUnitOfWork DomainInterfaces.IUnitOfWork
	iRedis      ApplicationInterfaces.IRedis
}

type sOrderHandlerCommandsParams struct {
	fx.In
	SConfig     *Config.SConfig
	IUnitOfWork DomainInterfaces.IUnitOfWork
	IRedis      ApplicationInterfaces.IRedis
}

func NewOrderHandlerCommands(sOrderHandlerCommandsParams sOrderHandlerCommandsParams) ApplicationInterfaces.IOrderHandlerCommands {
	return &SOrderHandlerCommands{
		sConfig:     sOrderHandlerCommandsParams.SConfig,
		iUnitOfWork: sOrderHandlerCommandsParams.IUnitOfWork,
		iRedis:      sOrderHandlerCommandsParams.IRedis,
	}
}
