package Application

import (
	"go.uber.org/fx"
)

var Modules []fx.Option

var Module = fx.Module("Application", Modules...)
