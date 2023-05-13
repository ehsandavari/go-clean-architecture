package config

type Grpc struct {
	IsEnabled     bool
	Port          string `validate:"required"`
	IsDevelopment bool
}
