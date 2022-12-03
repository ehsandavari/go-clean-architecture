package Commands

import (
	"GolangCodeBase/Application/Common"
	"GolangCodeBase/Domain/Entities"
	"GolangCodeBase/Domain/Interfaces"
	"context"
	"fmt"
)

func (r *SOrderHandlerCommands) PublishOrderCommand(ctx context.Context, orderEntity Entities.OrderEntity) {
	err := r.iUnitOfWork.Do(func(work Interfaces.IUnitOfWork) error {
		work.OrderRepository().Add(orderEntity)
		orderEntity.Id += 1
		work.OrderRepository().Add(orderEntity)
		return nil
	})
	fmt.Println(err)
	err = r.iRedis.Publish(ctx, r.sConfig.Redis.Queues["Orders"], Common.MarshalJson(orderEntity))
	if err != nil {
		panic(err)
	}
}
