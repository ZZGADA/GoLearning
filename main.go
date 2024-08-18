package main

import (
	"basicLearning/src/config"
	"basicLearning/src/connectToMysql"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"path"
	"reflect"
	"runtime"
)

var LogConf = struct {
	Dir     string `yaml:"dir"`
	Name    string `yaml:"name"`
	Level   string `yaml:"level"`
	MaxSize int    `yaml:"max_size"`
}{
	Dir:     "./logs",
	Name:    "yourlogname.log",
	Level:   "trace",
	MaxSize: 100,
}

func nacosTemp(nacosConfigCenter config_client.IConfigClient) {
	val, err := nacosConfigCenter.GetConfig(vo.ConfigParam{
		DataId: "basic-learning-go.yaml",
		Group:  "GO",
	})
	fmt.Println("hello world")
	fmt.Println(val, err)
}

func mysqlTemp(nacosConfigCenter config_client.IConfigClient) {
	// 定义一个切片
	var users []connectToMysql.TransactionStudy
	connectToMysql.ConnectToMysql(nacosConfigCenter)
	// 查询主键 in （2，4，5）中的
	connectToMysql.MysqlClient.Find(&users, []int{2, 4, 5})
	fmt.Println(reflect.TypeOf(users).Kind())
	for _, val := range users {
		fmt.Println(val)
	}

	//var user1 connectToMysql.TransactionStudy
	connectToMysql.MysqlClient.Find(&users, "age > ? and address = ?", 22, "HN")
	for _, val := range users {
		fmt.Printf("%#v", val)
	}

}

func getProjectRoot() {
	// 通过Caller获取当前执行文件的目录地址
	// 然后获取父级目录
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(filename)
	fmt.Println(root)
}

func main() {
	nacosConfig := config.Load()
	fmt.Println(nacosConfig)

}
