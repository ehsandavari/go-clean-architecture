package config

import (
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/postgres"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/redis"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"log"
)

type SConfig struct {
	Service  *SService         `validate:"required"`
	Postgres *postgres.SConfig `validate:"required"`
	Redis    *redis.SConfig    `validate:"required"`
}

func NewConfig() (*SConfig, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

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
