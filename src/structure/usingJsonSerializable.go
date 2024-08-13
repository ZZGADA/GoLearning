package structure

import (
	"encoding/json"
	"fmt"
)

// Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的
//`key1:"value1" key2:"value2"`

type house struct {
	Size  int    // 字段名首字母必须大写，否则json包访问不到字段数据
	Style string `json:"style"`
}

func UsingJsonSerialize() {
	h1 := house{
		100,
		"中国风",
	}
	data, err := json.Marshal(h1) // 序列化函数
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	// 没有Tag的情况下 全是大写的
	fmt.Printf("%s\n", data) // {"Size":100,"Style":"中国风"}

	str := `{"Size":100,"style":"中国风","hh":"oo"}` // JSON格式的字符串  json字段多了没事
	var h2 house
	// 反序列化函数，参数是[]byte类型的JSON格式字符串和结构体指针（因为要对结构体变量进行修改）
	// 先转换成字节数组
	err = json.Unmarshal([]byte(str), &h2)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", h2)

}
