package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Load() *NacosConfig {

	// 配置读取yaml 文件
	viper.SetConfigName("application.yaml") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")             // 或viper.SetConfigType("YAML")
	viper.AddConfigPath("./")               // 配置文件路径
	err := viper.ReadInConfig()             // 查找并读取配置文件
	if err != nil {                         // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	var nacosConfig NacosConfig
	if err := viper.UnmarshalKey("nacos", &nacosConfig); err != nil {
		panic("viper 转换对象错误 ")
	}

	return &nacosConfig
}
