package Commands

import (
	DomainEntities "GolangCodeBase/Domain/Entities"
	"context"
	"encoding/json"
	"fmt"
)

func (r SOrderHandlerCommands) SubscribeOrderCommand(ctx context.Context) error {
	channel := r.iRedis.Subscribe(ctx, r.sConfig.Redis.Queues["Orders"])
	go func() {
		orderEntity := DomainEntities.OrderEntity{}
		for {
			select {
			case channelData := <-channel:
				if err := json.Unmarshal([]byte(channelData), &orderEntity); err != nil {
					panic(err)
				}
				add := r.iUnitOfWork.OrderRepository().Add(orderEntity)
				fmt.Println(add)
			}
		}
	}()
	return nil
}
