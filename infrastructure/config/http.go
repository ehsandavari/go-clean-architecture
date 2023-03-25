package config

type Http struct {
	Mode string `validate:"required"`
	Host string `validate:"required"`
	Port string `validate:"required"`
}
