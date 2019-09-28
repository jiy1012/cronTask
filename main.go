package main

import (
	"cronTask/config"
	"flag"
)

func main() {
	cFile := flag.String("-conf", "./conf/cron.yaml", "配置文件地址")

	config.InitConfig(*cFile)
	config.ReloadCr()
	config.Cr.Start()
	select {}
}
