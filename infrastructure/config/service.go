package config

import (
	"github.com/ehsandavari/go-clean-architecture/domain/constants"
)

type SService struct {
	Id                     int                   `validate:"required"`
	Name                   string                `validate:"required"`
	Version                string                `validate:"required"`
	Mode                   constants.ServiceMode `validate:"required"`
	GracefulShutdownSecond byte                  `validate:"required"`
	Api                    *SApi                 `validate:"required"`
	Grpc                   *Grpc                 `validate:"required"`
}
