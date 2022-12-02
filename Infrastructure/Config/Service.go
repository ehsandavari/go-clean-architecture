package Config

import (
	"GolangCodeBase/Infrastructure/Logger"
)

type SService struct {
	Logger Logger.LogConfig `mapstructure:"Logger"`
}