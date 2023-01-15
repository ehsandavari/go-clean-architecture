package config

import (
	"github.com/ehsandavari/golang-clean-architecture/infrastructure"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/logger"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/postgres"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/redis"
	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
)

func init() {
	infrastructure.Modules = append(infrastructure.Modules,
		fx.Provide(NewConfig),
		fx.Provide(func(config *SConfig) (logger.SConfig, postgres.SConfig, redis.SConfig) {
			return config.Service.Logger, config.Postgres, config.Redis
		}),
	)
}

type SConfig struct {
	Service  SService         `validate:"required"`
	Postgres postgres.SConfig `validate:"required"`
	Redis    redis.SConfig    `validate:"required"`
}

func NewConfig() (*SConfig, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := new(SConfig)
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	if err := validator.New().Struct(config); err != nil {
		log.Fatalln(err.Error())
	}
	return config, nil
}

func (r SConfig) GetServiceId() uint16 {
	return r.Service.Id
}

func (r SConfig) GetServiceName() string {
	return r.Service.Name
}

func (r SConfig) GetServiceVersion() string {
	return r.Service.Version
}
