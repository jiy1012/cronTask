package config

type Config struct {
	Tasks []TaskConfig `mapstructure:"tasks"`
}

type TaskConfig struct {
	Schedule string
	Message  string
	AtAll    bool //是否@所有人
	AtPhone  string //@某个人
	Dingding string
}
