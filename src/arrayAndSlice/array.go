package arrayAndSlice

import (
	"fmt"
	"reflect"
	"unsafe"
)

func ArrayLearning() int {
	var intSize int32
	fmt.Println("the size of int in M2 chip is ", unsafe.Sizeof(intSize), "bytes") // 在 M2 芯片上 8 bytes

	// 初始化数组
	var numbersHashNumber = [5]int{1, 2, 3, 4, 5}
	numbersNull := [5]int{} // 创建一个长度为 5 的空数组，所有元素初始化为 0 这个是直接初始化
	var numberSign [5]int   // 创建一个长度为 5 的空数组，所有元素初始化为 0 这个是声明一个数组 然后默认初始化 这种是显示声明
	slice1 := []int{}       //空切片声明
	var slice2 = []int{}    // 空切片声明
	var slice3 []int
	println(slice3)

	// 循环数组
	fmt.Println("遍历开始")
	for index, val := range numbersHashNumber {
		s := fmt.Sprintf(" array index is %d, val is %d", index, val)
		fmt.Println(s)
		numbersNull[index] = val * index
		//numberSign = append(slice, numbersHashNumber[index])
	}
	fmt.Println("遍历结束")

	// 查看创建的空数组
	fmt.Println("array is ", numbersNull, len(numbersNull), cap(numbersNull), reflect.TypeOf(numbersNull).Kind())
	fmt.Println("array is ", numberSign, len(numberSign), cap(numberSign), reflect.TypeOf(numberSign).Kind())

	// 查看空切片
	fmt.Println("empty slice ", slice1, len(slice1), cap(slice1), reflect.TypeOf(slice1).Kind())
	fmt.Println("empty slice ", slice2, len(slice2), cap(slice2), reflect.TypeOf(slice2).Kind())
	return 1
}
