package redis

type SConfig struct {
	URL    string            `validate:"required"`
	Queues map[string]string `validate:"required"`
}
