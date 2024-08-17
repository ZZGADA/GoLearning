package main

import (
	"basicLearning/src/config"
	ginstudy "basicLearning/src/ginStudy"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func temp() {

}

func main() {
	nacosConfig := config.Load()
	nacosConfigCenter := ginstudy.InitNacosConfigCenter(nacosConfig)
	val, err := nacosConfigCenter.GetConfig(vo.ConfigParam{
		DataId: "basic-learning-go.yaml",
		Group:  "GO",
	})
	fmt.Println("hello world")
	fmt.Println(val, err)
}
