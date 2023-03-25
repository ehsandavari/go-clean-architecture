package config

import (
	"github.com/ehsandavari/golang-clean-architecture/infrastructure/logger"
)

type SService struct {
	Id      uint16          `validate:"required"`
	Name    string          `validate:"required"`
	Version string          `validate:"required"`
	Http    Http            `validate:"required"`
	Logger  *logger.SConfig `validate:"required"`
}
