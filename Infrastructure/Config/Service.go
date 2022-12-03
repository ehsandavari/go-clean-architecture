package Config

import (
	"GolangCodeBase/Infrastructure/Logger"
)

type SService struct {
	Id      uint16         `mapstructure:"id"`
	Name    string         `mapstructure:"name"`
	Version string         `mapstructure:"version"`
	Logger  Logger.SConfig `mapstructure:"logger"`
	Grpc    Logger.SConfig `mapstructure:"logger"`
}
