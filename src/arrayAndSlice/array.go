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

	// 对于没有初始化长度的数组就是切片 slice
	slice1 := []int{}            //空切片声明
	var slice2 = []int{}         // 空切片声明
	var slice3 []int             // 使用这种形式的切片最好
	slice4 := make([]int, 5, 12) //初始化元素有5个 初始值都为0 但是切片的容量是12

	// 循环数组
	fmt.Println("遍历开始")
	// 当前写法 是遍历出索引的具体值
	for index, val := range numbersHashNumber {
		// _, num :=range numberHashNumber 那么就不需要index 直接只获取value值
		s := fmt.Sprintf(" array index is %d, val is %d", index, val)
		fmt.Println(s)
		numbersNull[index] = val * index
		slice3 = append(slice3, numbersHashNumber[index])
	}
	fmt.Println("遍历结束")

	// 查看创建的空数组
	fmt.Println("array is ", numbersNull, len(numbersNull), cap(numbersNull), reflect.TypeOf(numbersNull).Kind())
	fmt.Println("array is ", numberSign, len(numberSign), cap(numberSign), reflect.TypeOf(numberSign).Kind())

	// 查看空切片
	fmt.Println("empty slice1 ", slice1, len(slice1), cap(slice1), reflect.TypeOf(slice1).Kind())
	fmt.Println("empty slice2 ", slice2, len(slice2), cap(slice2), reflect.TypeOf(slice2).Kind())
	fmt.Println("empty slice3 ", slice3, len(slice3), cap(slice3), reflect.TypeOf(slice3).Kind())
	fmt.Println("empty slice4 ", slice4, len(slice4), cap(slice4), reflect.TypeOf(slice4).Kind())
	slice4 = append(slice4, 1) //向元素末尾追加一个元素 元素长度 len扩充到6 slice容量不变
	fmt.Println("empty slice4 ", slice4, len(slice4), cap(slice4), reflect.TypeOf(slice4).Kind())
	slice4 = append(slice4, 1, 2, 3, 4, 5, 6, 7) //向元素末尾追加元素 同时扩容为原来的两倍
	fmt.Println("empty slice4 ", slice4, len(slice4), cap(slice4), reflect.TypeOf(slice4).Kind())

	slice5 := numbersNull[2:len(numbersNull)] // 切片是数组的引用  元素修改则改变
	slice5[0] = 100
	slice5 = append(slice5, -1) // 但是元素增加并不影响原数组
	fmt.Println("empty slice5 ", slice5, len(slice5), cap(slice5), reflect.TypeOf(slice5).Kind())
	fmt.Println("array is ", numbersNull, len(numbersNull), cap(numbersNull), reflect.TypeOf(numbersNull).Kind())

	return 1
}
