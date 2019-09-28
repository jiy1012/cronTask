package config

import "github.com/robfig/cron/v3"

type Config struct {
	Tasks []TaskConfig `mapstructure:"tasks"`
}

type TaskConfig struct {
	Title    string
	Schedule string
	Message  string
	AtAll    bool   //是否@所有人
	AtPhone  string //@某个人
	Dingding string
}

type Tm struct {
	Md5 string
	Id  cron.EntryID
}

type SchduleActiveCallback func(*TaskConfig)

type ScheduleJob struct {
	task     *TaskConfig
	callback SchduleActiveCallback
}
