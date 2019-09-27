package main

import (
	"cronTask/comm"
	"cronTask/config"
	"flag"
	"fmt"
	"github.com/robfig/cron"
	"strings"
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
		t := task
		fmt.Println(t)
		job := NewScheduleJob(&t, sendDingDing)
		cr.AddJob(task.Schedule, job)
	}
	cr.Start()
	select {}

}

func sendDingDing(task *config.TaskConfig) {
	dd := comm.DDRobotStruct{}
	dd.Text.Content = task.Message
	dd.Msgtype = "string"
	dd.At.Atmobiles = strings.Split(task.AtPhone, ",")
	dd.At.Isatall = task.AtAll
	fmt.Printf("%+v\n", dd)
	comm.SendDDRobot(dd, task.Dingding)
}
