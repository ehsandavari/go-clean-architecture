package Config

import (
	"GolangCodeBase/Infrastructure/Postgres"
	"GolangCodeBase/Infrastructure/Redis"
	"github.com/spf13/viper"
	"os"
)

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
