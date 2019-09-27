package main

import (
	"cronTask/config"
	"flag"
	"fmt"
	"github.com/robfig/cron"
)

type SchduleActiveCallback func(*config.TaskConfig)

type ScheduleJob struct {
	task     *config.TaskConfig
	callback SchduleActiveCallback
}

func (this *ScheduleJob) Run() {
	if nil != this.callback {
		this.callback(this.task)
	} else {
		fmt.Println("error no callback")
	}
}

func NewScheduleJob(task *config.TaskConfig, callback SchduleActiveCallback) *ScheduleJob {
	instance := &ScheduleJob{
		task:     task,
		callback: callback,
	}
	return instance
}

func main() {
	cFile := flag.String("-conf", "./conf/cron.yaml", "配置文件地址")
	var c config.Config
	config.InitConfig(*cFile, &c)

	cr := cron.New()
	for _, task := range c.Tasks {
		job := NewScheduleJob(&task, sendDingDing)
		cr.AddJob(task.Schedule, job)
	}
	cr.Start()
	select {}

}

func sendDingDing(task *config.TaskConfig) {
	fmt.Println("dingding:", task.Dingding, "message:", task.Message)
}
