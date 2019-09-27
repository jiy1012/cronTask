package config

type Config struct {
	Tasks []TaskConfig `mapstructure:"tasks"`
}

type TaskConfig struct {
	Schedule string
	Message  string
	Dingding string
}
