package redis

import (
	"context"
	ApplicationInterfaces "github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure"
	"github.com/go-redis/redis/v9"
	"go.uber.org/fx"
	"log"
)

func init() {
	infrastructure.Modules = append(infrastructure.Modules, fx.Provide(NewRedis))
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
