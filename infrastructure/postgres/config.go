package postgres

type SConfig struct {
	URL string `validate:"required"`
}
