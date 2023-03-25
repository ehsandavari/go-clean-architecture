package postgres

type SConfig struct {
	Host         string `validate:"required"`
	Port         string `validate:"required"`
	User         string `validate:"required"`
	Password     string `validate:"required"`
	DatabaseName string `validate:"required"`
	SslMode      string `validate:"required"`
	TimeZone     string `validate:"required"`
}
