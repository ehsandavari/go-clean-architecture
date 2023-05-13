package config

type STracer struct {
	IsEnabled bool
	Host      string `validate:"required"`
	Port      string `validate:"required"`
	Sampler   bool
	UseStdout bool
}
