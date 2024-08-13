package structure

import (
	lambdafunction "basicLearning/src/lambdaFunction"
	"fmt"
	"reflect"
)

// 接口一旦定义 实现的结构体必须全部实现其中的方法
type Phone interface {
	call()
	message()
}

// NokiaPhone 定义了结构体NokiaPhone
type NokiaPhone struct {
}

// 实现接口方法
// 声明是哪个结构体使用
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

func (nokiaPhone NokiaPhone) message() {
	fmt.Printf("message 调用后 变量结构 %#v , 地址是：%p\n", nokiaPhone, &nokiaPhone)
	fmt.Println("nokia is messaging you")
}

// IPhone 定义了结构体IPhone
type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func (iPhone *IPhone) message() {
	fmt.Println("IPhone is messaging you")
}

// TryUsingInterface // 用于外部调用
func TryUsingInterface() {
	// 定义Phone的变量
	var phoneN Phone
	var phoneI Phone
	var phone Phone
	var phoneNTest *NokiaPhone

	//new 出来的对象 返回的是对象的指针
	// new 一个NokiaPhone的结构体 多态实现
	phoneN = new(NokiaPhone)
	// 确定--->go 将pointer.func()的新式会转换成 (*pointer).func()的形式 可以反编译看看
	// CALL basicLearning/src/structure.(*NokiaPhone).call(SB)
	phoneN.call()

	// new 一个IPhone的结构体 多态实现
	phoneI = new(IPhone)
	phoneI.call()

	phone = new(IPhone)
	phone.call()

	phoneTest := new(NokiaPhone)
	phoneStructTest := NokiaPhone{} // 这种出来的是结构体
	phoneTest.call()

	phoneNTest = new(NokiaPhone)

	// reflect.TypeOf的结果是 ptr --->指针
	fmt.Println("打印类型 phoneN", reflect.TypeOf(phoneN).Kind(), "变量类型是", reflect.TypeOf(phoneN).Kind())
	fmt.Println("打印类型 phoneI", reflect.TypeOf(phoneI).Kind())
	fmt.Println("打印类型 phone", reflect.TypeOf(phone).Kind())
	lambdafunction.UsingCodeBlock("-")

	fmt.Println("打印类型 phoneTest", reflect.TypeOf(phoneTest).Kind(), "变量类型是", reflect.TypeOf(*phoneTest).Kind())
	fmt.Println("打印类型 phoneTest", reflect.TypeOf(phoneTest).Kind(), "变量类型是", reflect.TypeOf(&phoneTest).Kind())
	fmt.Printf("phoneTest:对象的值%v\n", phoneTest)  // %v打印出来的是对象的值
	fmt.Printf("phoneTest:对象的值%v\n", *phoneTest) // %v打印出来的是对象的值
	fmt.Printf("phoneTest:地址%p\n", phoneTest)    // %v打印出来的是对象的值
	fmt.Printf("phoneTest:二级地址%v\n", &phoneTest)
	lambdafunction.UsingCodeBlock("-")

	fmt.Println("打印类型 phoneNTest", reflect.TypeOf(phoneNTest).Kind(), "变量类型是", reflect.TypeOf(*phoneNTest).Kind())
	fmt.Printf("phoneNTest:对象的值%v\n", phoneNTest)
	fmt.Printf("phoneNTest:地址%p\n", phoneNTest)
	fmt.Printf("phoneNTest:二级地址%v\n", &phoneNTest)
	lambdafunction.UsingCodeBlock("-")

	fmt.Printf("phoneI:对象的值%v\n", phoneI)
	fmt.Printf("phoneI:地址%p\n", phoneI)
	fmt.Printf("phoneI:二级地址%v\n", &phoneI)
	lambdafunction.UsingCodeBlock("-")

	fmt.Printf("phoneStructTest:对象的值%v\n", phoneStructTest)
	fmt.Printf("phoneStructTest:地址%p\n", &phoneStructTest)
	fmt.Println("打印类型 phoneStructTest", reflect.TypeOf(phoneStructTest).Kind(), "变量类型是", reflect.TypeOf(&phoneStructTest).Kind())
	lambdafunction.UsingCodeBlock("+")

	fmt.Printf("message 调用前 变量结构 %#v , 地址是：%p\n", phoneN, phoneN)
	phoneN.message()
	lambdafunction.UsingCodeBlock("+")

}
