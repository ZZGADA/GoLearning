package errorcatch

import (
	lambdafunction "basicLearning/src/lambdaFunction"
	"errors"
	"fmt"
)

// 一个可能返回错误的函数
func divide(a, b int) (int, error) {
	// 自定义错误条件
	if b == 0 {
		// 自定义直接返回错误条件
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func TryDivideError(chu int) {
	// 使用带初始化的 if 语句来处理错误
	// 即divide 返回两个结果 如果 error != nil则表示有报错
	if result, err := divide(10, chu); err != nil {
		fmt.Println("Error:", err)
	} else {
		// 否则逻辑正确
		fmt.Println("Result:", result)
	}
}

type DivideError struct {
	dividee int
	divider int
}

// 实现 `error` 接口
// 结构体DivideError实现error接口的Error方法
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    divide: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// Divide // 定义 `int` 类型除法运算的函数
func Divide(varDivide int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDivide,
			divider: varDivider,
		}
		errorMsg = dData.Error() // 调用Error方法
		return
	} else {
		return varDivide / varDivider, ""
	}

}

// TrySelfDefineError //
func TrySelfDefineError() {
	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

	lambdafunction.UsingCodeBlock("*")
	simpleTry()
	fmt.Println("simpleTry执行结束")
}

// 在这个使用场景下 当除数是0的时候 发生异常 然后recover拦截 程序继续执行 没有中断
func simpleTry() {
	// 函数退出的时候执行
	defer func() {
		error := recover()
		// 函数推出就会执行 但是只有接收到异常才会有下一步操作
		if error != nil {
			fmt.Println(error, "amazing")
		} else {
			fmt.Println("defer 最后处理 recover没有拦截到异常")
			//panic("panic 严重异常报错")  //panic直接中断
		}
	}()
	ko := 2 % 3
	a := 10 / ko
	fmt.Println(a)
	fmt.Println("代码块到底了")

}
