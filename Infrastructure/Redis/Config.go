package Redis

type SConfig struct {
	URL    string            `mapstructure:"url"`
	Queues map[string]string `mapstructure:"queues"`
}