package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Bases Base

var v = viper.New()

func InitConfig() {

	v.SetConfigFile("./config.yaml") // 指定配置文件路径
	v.SetConfigName("config")        // 配置文件名称(无扩展名)
	v.SetConfigType("yaml")          // 如果配置文件的名称中没有扩展名，则需要配置此项
	v.AddConfigPath("./")            // 还可以在工作目录中查找配置

	err := v.ReadInConfig() // 查找并读取配置文件
	if err != nil {         // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 监控配置文件变化
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&Bases); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&Bases); err != nil {
		fmt.Println(err)
	}
}

func Get(key string) interface{} {
	return v.Get(key)
}
