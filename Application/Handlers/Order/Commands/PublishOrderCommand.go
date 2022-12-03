package Commands

import (
	"GolangCodeBase/Application/Common"
	"GolangCodeBase/Domain/Entities"
	"context"
)

func (r *SOrderHandlerCommands) PublishOrderCommand(ctx context.Context, orderEntity Entities.OrderEntity) {
	r.iLogger.Info("sadkmaksdlasdj")
	err := r.iRedis.Publish(ctx, r.sConfig.Redis.Queues["Orders"], Common.MarshalJson(orderEntity))
	if err != nil {
		panic(err)
	}
}
