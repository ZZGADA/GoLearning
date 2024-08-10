package lambdafunction

import (
	"fmt"
	"strings"
)

// 定义全局变量
var separateCodeBlockMethod int32 = 0

// SeparateCodeBlock /*Separate CodeBlock 输出分割符号 有三种实现方法 ===>字符串拼接的三种方法*/
func SeparateCodeBlock(separateVal string) string {
	var separateString string
	switch separateCodeBlockMethod % 3 {
	case 0:
		// 字符串拼接的方式 比较慢
		separateString = func() string {
			s := ""
			for i := 0; i < 50; i++ {
				s += separateVal
			}
			return "字符串拼接方法 ->\t" + s
		}()
	case 1:
		// 使用stringBuilder的形式创建 速度最快
		// strings.Builder 是 Go 语言的 strings 包中提供的一个类型，用于高效地构建字符串。
		// 它是专门为减少内存分配和复制开销而设计的，适合需要频繁拼接字符串的场景。
		separateString = func() string {
			var builder = strings.Builder{}
			for i := 0; i < 50; i++ {
				builder.WriteString(separateVal)
			}
			return "strings.Builder方法 ->\t" + builder.String()
		}()
	case 2:
		// 使用string.join方法 拼接字符串 速度比strings.Builder差 但是比第一中好
		separateString = func() string {
			arr := make([]string, 50)
			for i := range arr {
				arr[i] = separateVal
			}
			return "数组join方法 ->\t" + strings.Join(arr, "")
		}()

	}
	// mention Go 语言不允许在同一行中直接结合 ++ 和 % 操作
	//(separateCodeBlockMethod++)%=3
	separateCodeBlockMethod++
	separateCodeBlockMethod %= 3

	return separateString
}

// CallFunctionOfSeparateCodeBlock 定义一个回掉函数类型
type CallFunctionOfSeparateCodeBlock func(value string) string

// 定义一个接受回调函数的函数 参数是函数指针
func processSeparateBlock(separateVal string, callBack CallFunctionOfSeparateCodeBlock) {
	// 实现处理逻辑 ，和调用回调函数
	methodNum := separateCodeBlockMethod + 1
	sepBlock := callBack(separateVal)
	fmt.Println(fmt.Sprintf("调用第%d种分割方法\t %s", methodNum, sepBlock))
}

// 封装方法
func usingCodeBlock(separateVal string) {
	// 使用回调函数
	processSeparateBlock(separateVal, SeparateCodeBlock)
}

// UsingLambda /*Using lambda 使用匿名函数*/
func UsingLambda() {
	fmt.Println("using lambda ---> ")

	// simple using lambda

	usingCodeBlock("=")

	// 定义并立即调用匿名函数
	//result := func(x, y int) int {
	//	return x + y
	//}(3, 4) // 传递参数

	//fmt.Println("Result:", result) // 输出: Result: 7
}
