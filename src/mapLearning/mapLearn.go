package maplearning

import "fmt"

// TryMap /**/
func TryMap() {
	// 创建一个空的 map，key 是 int 类型，value 是 float32 类型
	// 创建一个初始容量为 10 的 Map
	//m := make(map[string]int, 10)
	map1 := make(map[int32]float64)

	// 向 map1 中添加 key-value 对
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	// 遍历 map1，读取 key 和 value
	for key, value := range map1 {
		// 打印 key 和 value
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	// 遍历 map1，只读取 key
	for key := range map1 {
		// 打印 key
		fmt.Printf("key is: %d\n", key)
	}

	// 遍历 map1，只读取 value
	for _, value := range map1 {
		// 打印 value
		fmt.Printf("value is: %f\n", value)
	}

	val, ifHasKey := map1[100]
	if ifHasKey {
		fmt.Println("has key 100 and the value is ", val)
	} else {
		fmt.Println("key not exist")
	}

	// delete(countryCapitalMap, "France") 通过这种形式删除元素

}
