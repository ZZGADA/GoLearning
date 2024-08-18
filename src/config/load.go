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

func logTest() {
	//nacosConfigCenter := ginstudy.InitNacosConfigCenter(nacosConfig)
	////nacosTemp(nacosConfigCenter)
	//mysqlTemp(nacosConfigCenter)
	//config.LogConfig("debug")

	// 设置日志格式。
	//logrus.SetFormatter(&logrus.TextFormatter{
	//	TimestampFormat: "2006-01-02 15:04:05.000",
	//})
	//switch "debug" {
	//case "trace":
	//	logrus.SetLevel(logrus.TraceLevel)
	//case "debug":
	//	logrus.SetLevel(logrus.DebugLevel)
	//case "info":
	//	logrus.SetLevel(logrus.InfoLevel)
	//case "warn":
	//	logrus.SetLevel(logrus.WarnLevel)
	//case "error":
	//	logrus.SetLevel(logrus.ErrorLevel)
	//case "fatal":
	//	logrus.SetLevel(logrus.FatalLevel)
	//case "panic":
	//	logrus.SetLevel(logrus.PanicLevel)
	//}
	//logrus.SetReportCaller(true) // 打印文件、行号和主调函数。
	//
	//// 实现日志滚动。
	//// Refer to https://www.cnblogs.com/jssyjam/p/11845475.html.
	//logger := &lumberjack.Logger{
	//	Filename:   fmt.Sprintf("%v/%v", LogConf.Dir, LogConf.Name), // 日志输出文件路径。
	//	MaxSize:    LogConf.MaxSize,                                 // 日志文件最大 size(MB)，缺省 100MB。
	//	MaxBackups: 10,                                              // 最大过期日志保留的个数。
	//	MaxAge:     30,                                              // 保留过期文件的最大时间间隔，单位是天。
	//	LocalTime:  true,                                            // 是否使用本地时间来命名备份的日志。
	//}
	//logrus.SetOutput(logger)
	//
	//// 日志示例
	//logrus.Info("This is an info message")
	//logrus.Warn("This is a warning message")
	//logrus.Error("This is an error message")

	//var mu sync.Mutex
	//
	//// 启动一个 goroutine，获取互斥锁并解锁
	//go func() {
	//	mu.Lock()
	//	// 解锁
	//	mu.Unlock()
	//}()
	//
	//// 主 goroutine 等待一段时间，以确保 goroutine 能够获取并解锁互斥锁
	//// 这只是一个示例，实际应用中应使用更可靠的同步机制
	//// time.Sleep(time.Second)
	//
	//// 主 goroutine 获取互斥锁
	//mu.Lock()
	//// 解锁
	//mu.Unlock()
}
