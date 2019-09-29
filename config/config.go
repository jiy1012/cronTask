package config

import (
	"cronTask/comm"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

var TaskMap map[string]*Tm = map[string]*Tm{}
var Cr = cron.New()
var C = Config{}

func InitConfig(cFile string) {
	dir := filepath.Dir(cFile)
	ext := strings.TrimLeft(filepath.Ext(cFile), ".")
	file := strings.TrimRight(filepath.Base(cFile), "."+ext)
	viper.SetConfigName(file) //把json文件换成yaml文件，只需要配置文件名 (不带后缀)即可
	viper.AddConfigPath(dir)  //添加配置文件所在的路径
	viper.SetConfigType(ext)  //设置配置文件类型
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
	err = viper.Unmarshal(&C)
	if err != nil {
		fmt.Printf("config file unmarshal error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(C)
	viper.WatchConfig() //监听配置变化
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置发生变更：", e.Name)
		err = viper.Unmarshal(&C)
		if err != nil {
			fmt.Printf("config file unmarshal error: %s\n", err)
			os.Exit(1)
		}
		fmt.Println(C)
		ReloadCr()
	})

	return
}

//func beginCr() {
//	for _, task := range C.Tasks {
//		t := task
//		fmt.Println(t)
//		job := schedule.NewScheduleJob(&t, comm.SendTaskDingDing)
//		id, _ := Cr.AddJob(task.Schedule, job)
//		tm := &Tm{
//			Md5: comm.MD5Hash(comm.JsonEncode(task)),
//			Id:  id,
//		}
//		if _, ok := TaskMap[comm.MD5Hash(task.Title)]; !ok {
//			TaskMap[comm.MD5Hash(task.Title)] = tm
//		}
//	}
//}
func ReloadCr() {
	//备份旧的keyMap
	keyMap := map[string]int{}
	//旧的taskMap
	for key := range TaskMap {
		keyMap[key] = 1
	}
	//新的task配置
	for _, task := range C.Tasks {
		key := comm.MD5Hash(task.Title)
		//新配置启动过
		if tTm, ok := TaskMap[key]; ok {
			//删除已经启动过的key
			delete(keyMap, key)
			//如果新配置和旧配置不一样 reload 一样则忽略
			if tTm.Md5 != comm.MD5Hash(comm.JsonEncode(task)) {
				Cr.Remove(tTm.Id)
				t := task
				fmt.Printf("reload:%+v \n", t)
				job := NewScheduleJob(&t, SendTaskDingDing)
				id, err := Cr.AddJob(task.Schedule, job)
				if err != nil {
					fmt.Println("add job err:", err.Error())
				}
				TaskMap[key].Id = id
				TaskMap[key].Md5 = comm.MD5Hash(comm.JsonEncode(task))
			}
		} else {
			//新配置没启动过，启动
			t := task
			fmt.Printf("start:%+v \n", t)
			job := NewScheduleJob(&t, SendTaskDingDing)
			id, err := Cr.AddJob(task.Schedule, job)
			if err != nil {
				fmt.Println("add job err:", err.Error())
			}
			tm := &Tm{
				Md5: comm.MD5Hash(comm.JsonEncode(task)),
				Id:  id,
			}
			if _, ok := TaskMap[comm.MD5Hash(task.Title)]; !ok {
				TaskMap[comm.MD5Hash(task.Title)] = tm
			}
		}

	}
	//剩余要删除的任务
	for key := range keyMap {
		if v, ok := TaskMap[key]; ok {
			Cr.Remove(v.Id)
			fmt.Printf("remove task:%+v \n", v)
			delete(TaskMap, key)
		}
	}
}

func (this *ScheduleJob) Run() {
	if nil != this.callback {
		this.callback(this.task)
	} else {
		fmt.Println("error no callback")
	}
}

func NewScheduleJob(task *TaskConfig, callback SchduleActiveCallback) *ScheduleJob {
	instance := &ScheduleJob{
		task:     task,
		callback: callback,
	}
	return instance
}

func SendTaskDingDing(task *TaskConfig) {
	dd := comm.DDRobotStruct{}
	dd.Text.Content = task.Message
	dd.Msgtype = "string"
	dd.At.Atmobiles = strings.Split(task.AtPhone, ",")
	dd.At.Isatall = task.AtAll
	fmt.Printf("send task %+v\n", dd)
	comm.SendDDRobot(dd, task.Dingding)
}
