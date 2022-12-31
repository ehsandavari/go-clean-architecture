package config

import (
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/logger"
)

type SService struct {
	Id      uint16         `mapstructure:"id"`
	Name    string         `mapstructure:"name"`
	Version string         `mapstructure:"version"`
	Logger  logger.SConfig `mapstructure:"logger"`
	Grpc    logger.SConfig `mapstructure:"logger"`
}
