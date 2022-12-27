package Redis

import (
	ApplicationInterfaces "GolangCodeBase/Application/Common/Interfaces"
	"GolangCodeBase/Infrastructure"
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
	"log"
)

func init() {
	Infrastructure.Modules = append(Infrastructure.Modules, fx.Provide(NewRedis))
}

type SRedis struct {
	Client *redis.Client
}

func NewRedis(lc fx.Lifecycle, config SConfig) ApplicationInterfaces.IRedis {
	sRedis := new(SRedis)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			sRedis.Client = redis.NewClient(&redis.Options{
				Addr: config.URL,
			})
			log.Println("redis connection opened")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("redis connection closed")
			return sRedis.close()
		},
	})
	return sRedis
}

func (r SRedis) Publish(ctx context.Context, channelName string, message interface{}) error {
	return r.Client.Publish(ctx, channelName, message).Err()
}

func (r *SRedis) Subscribe(ctx context.Context, channelName string) <-chan string {
	var subscribe = r.Client.Subscribe(ctx, channelName)
	channel := make(chan string)
	go func() {
		for msg := range subscribe.Channel() {
			channel <- msg.Payload
		}
	}()
	return channel
}

func (r *SRedis) close() error {
	return r.Client.Close()
}
