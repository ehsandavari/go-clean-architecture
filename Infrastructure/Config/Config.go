package Config

import (
	"GolangCodeBase/Infrastructure/Postgres"
	"GolangCodeBase/Infrastructure/Redis"
	"github.com/spf13/viper"
	"os"
)

type SConfig struct {
	Service  SService         `mapstructure:"postgres"`
	Postgres Postgres.SConfig `mapstructure:"postgres"`
	Redis    Redis.SConfig    `mapstructure:"redis"`
}

func NewConfig() (*SConfig, error) {
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigName("Config.dev" + os.Getenv("env"))
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := new(SConfig)
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return config, nil
}
