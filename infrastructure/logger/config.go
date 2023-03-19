package logger

type SConfig struct {
	Level   string `validate:"required"`
	Mode    string `validate:"required"`
	Encoder string `validate:"required"`
}
