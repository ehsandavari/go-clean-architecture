package config

import (
	"github.com/ehsandavari/golang-clean-architecture/infrastructure"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/logger"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/postgres"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/redis"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"os"
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
	Service  SService         `mapstructure:"service"`
	Postgres postgres.SConfig `mapstructure:"postgres"`
	Redis    redis.SConfig    `mapstructure:"redis"`
}

func NewConfig() (*SConfig, error) {
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigName("config." + os.Getenv("env"))
	viper.SetConfigType("yml")
	viper.BindEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := new(SConfig)
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
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
