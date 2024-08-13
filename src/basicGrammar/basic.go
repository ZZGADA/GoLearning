package basicGrammar

import (
	"fmt"
	"strconv"
)

// BasicGrammar
func BasicGrammar() {
	str := "123" // 如果传入的不是整形 则报错
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str, num)
	}

	num2 := 123
	str2 := strconv.Itoa(num)
	fmt.Printf("整数 %d  转换为字符串为：'%s'\n", num2, str2)

	str3 := "12.3"
	f2, _ := strconv.ParseFloat(str3, 32) // _用于接受报错
	fmt.Printf("字符串 '%s' 转换为浮点数为：%f\n", str3, f2)
}
