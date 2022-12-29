package logger

type SConfig struct {
	LogLevel string `mapstructure:"level"`
	DevMode  bool   `mapstructure:"devMode"`
	Encoder  string `mapstructure:"encoder"`
}
