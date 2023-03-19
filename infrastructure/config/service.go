package config

import (
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/logger"
)

type SService struct {
	Id      uint16         `validate:"required"`
	Name    string         `validate:"required"`
	Version string         `validate:"required"`
	Logger  logger.SConfig `validate:"required"`
	//Grpc    logger.SConfig `validate:"required"`
}
