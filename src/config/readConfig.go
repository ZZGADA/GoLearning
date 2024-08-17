package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// ReadYaml //
func ReadYaml() {

	// 配置读取yaml 文件
	viper.SetConfigName("application.yaml") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")             // 或viper.SetConfigType("YAML")
	viper.AddConfigPath("./")               // 配置文件路径
	err := viper.ReadInConfig()             // 查找并读取配置文件
	if err != nil {                         // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	fmt.Println("Server Port:", viper.GetInt("server.port"))
	fmt.Println("Database User:", viper.GetStringMapString("databases.mysql"))
}

type ProjectSelf struct {
	Language  string   `yaml:"language"`
	Frame     string   `yaml:"frame"`
	Databases []string `yaml:"databases"`
	Mq        string   `yaml:"mq"`
	Es        string   `yaml:"es"`
}

type Project struct {
	Project ProjectSelf
}

//func ReadConfigFromNacos() {
//	val, _ := ginstudy.ReadNacosConfigYaml()
//	fmt.Println(val, reflect.TypeOf(val).Kind())
//	var tempMap Project
//
//	err := yaml.Unmarshal([]byte(val), &tempMap)
//	if err != nil {
//		panic("stop ")
//	}
//	fmt.Printf("%#v", tempMap)
//}

/*
mapstructure 标签的主要作用是指定结构体字段与映射中的键之间的对应关系。
这样，当你使用库（如 Viper）将映射数据解析到结构体时，库会根据 mapstructure 标签来匹配映射中的键和值，并将其填充到相应的结构体字段中。
*/
type NacosConfig struct {
	Dir                 Directory `yaml:"dir" mapstructure:"dir"`
	LogLevel            string    `yaml:"logLevel" mapstructure:"logLevel"`
	Username            string    `yaml:"username" mapstruture:"username"`
	Password            string    `yaml:"password" mapstruture:"password"`
	TimeoutMs           uint64    `yaml:"timeoutMs" mapstruture:"timeoutMs"`
	Ip                  string    `yaml:"ip" mapstruture:"ip"`
	Port                uint64    `yaml:"port" mapstruture:"port"`
	NotLoadCacheAtStart bool      `yaml:"notLoadCacheAtStart" mapstructure:"notLoadCacheAtStart"`
	ContextPath         string    `yaml:"contextPath" mapstructure:"contextPath"`
}
type Directory struct {
	Log   string `yaml:"log" mapstruture:"log"`
	Cache string `yaml:"cache" mapstruture:"cache"`
}
