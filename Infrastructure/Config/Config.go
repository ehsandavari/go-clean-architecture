package Config

import (
	"GolangCodeBase/Infrastructure"
	"GolangCodeBase/Infrastructure/Logger"
	"GolangCodeBase/Infrastructure/Postgres"
	"GolangCodeBase/Infrastructure/Redis"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"os"
)

func init() {
	Infrastructure.Modules = append(Infrastructure.Modules,
		fx.Provide(NewConfig),
		fx.Provide(func(config *SConfig) (Logger.SConfig, Postgres.SConfig, Redis.SConfig) {
			return config.Service.Logger, config.Postgres, config.Redis
		}),
	)
}

type SConfig struct {
	Service  SService         `mapstructure:"service"`
	Postgres Postgres.SConfig `mapstructure:"postgres"`
	Redis    Redis.SConfig    `mapstructure:"redis"`
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
