package config

import (
	"github.com/ehsandavari/go-clean-architecture/infrastructure/postgres"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"log"
)

type SConfig struct {
	Service  *SService         `validate:"required"`
	Postgres *postgres.SConfig `validate:"required"`
	Logger   *SLogger          `validate:"required"`
	Tracer   *STracer          `validate:"required"`
}

func NewConfig() *SConfig {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("error in read config ", err)
	}

	config := new(SConfig)
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalln("error in unmarshal config ", err)
	}

	if err := validator.New().Struct(config); err != nil {
		log.Fatalln("error in validate config ", err)
	}

	if !config.Service.Mode.IsValid() {
		log.Fatalln("service mode is not valid ", config.Service.Mode.String())
	}

	return config
}
