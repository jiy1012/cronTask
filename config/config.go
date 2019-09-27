package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

func InitConfig(cFile string, c *Config) {
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
	err = viper.Unmarshal(c)
	if err != nil {
		fmt.Printf("config file unmarshal error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(c)
	viper.WatchConfig() //监听配置变化
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置发生变更：", e.Name)
		err = viper.Unmarshal(c)
		if err != nil {
			fmt.Printf("config file unmarshal error: %s\n", err)
			os.Exit(1)
		}
		fmt.Println(c)
	})

	return
}
