package Interfaces

import (
	"GolangCodeBase/Domain/Entities"
	"context"
)

type (
	IOrderHandlerCommands interface {
		SubscribeOrderCommand(ctx context.Context) error
		PublishOrderCommand(ctx context.Context, orderEntity Entities.OrderEntity)
	}
)
