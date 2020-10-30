package config

import "github.com/robfig/cron/v3"

type Config struct {
	Tasks   []TaskConfig  `mapstructure:"tasks"`
	Holiday HolidayConfig `mapstructure:"holiday"`
}
type HolidayConfig struct {
	Url string
}
type TaskConfig struct {
	Title              string
	Schedule           string
	Message            string
	AtAll              bool   //是否@所有人
	AtPhone            string //@某个人
	Type               string
	WebhookUrl         string
	SkipChineseHoliday bool //是否跳过中国节假日
}

type Tm struct {
	Md5 string
	Id  cron.EntryID
}

type SchduleActiveCallback func(*Config, *TaskConfig)

type ScheduleJob struct {
	conf     *Config
	task     *TaskConfig
	callback SchduleActiveCallback
}
