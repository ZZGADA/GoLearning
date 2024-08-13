package structure

import (
	lambdafunction "basicLearning/src/lambdaFunction"
	"fmt"
)

type House struct {
	size, price float64
	style       string
}

// 返回结构体指针
func newHouse(size, price float64, style string) *House {
	return &House{
		size:  size,
		price: price,
		style: style,
	}
}

// 定义类的属性 具体的方法 用于执行逻辑操作
func (structure *House) increaseSize() {
	fmt.Printf("increaseSize 调用后 变量结构 %#v , 地址是：%p\n", structure, structure)
	structure.size++
}

func (structure House) decreaseSize() {
	// 如果方法的接收者是值类型，无论调用者是对象还是对象指针，修改的都是对象的副本，不影响调用者
	// 地址答打印可见 二者地址不同
	fmt.Printf("decreaseSize 调用后 变量结构 %#v , 地址是：%p\n", structure, &structure)
	structure.size--
}

// UsingStructureFunc // 使用构造函数
func UsingStructureFunc() {
	p1 := newHouse(100, 80, "中国风")
	fmt.Printf("increaseSize 调用前 变量结构 %#v , 地址是：%p\n", p1, p1)
	p1.increaseSize() // 执行完后 size有++操作了
	lambdafunction.UsingCodeBlock("+")
	fmt.Printf("decreaseSize 调用前 变量结构 %#v , 地址是：%p\n", p1, p1)
	fmt.Printf("decreaseSize 调用前 判断是否是二级指针 变量结构 %#v , 地址是：%p\n", p1, &p1)
	p1.decreaseSize()
	fmt.Printf("decreaseSize 调用final 变量结构 %#v , 地址是：%p\n", p1, p1)

	// %#v：输出值的 Go 语言语法表示。适用于显示数据结构的详细信息，包括字段名和类型。
	fmt.Printf("%#v\n", p1) // &main.House{size:100, price:80, style:"中国风"}

}

type Person struct {
	age int
}

func (p Person) howOld() int {
	temp := p.age
	//p.age++
	//fmt.Println("test", p.age)
	return temp
}

func (p *Person) growUp() {
	p.age += 1
}

// StructureValueAndPointer
func StructureValueAndPointer() {
	// qcrao 是值类型
	qcrao := Person{age: 18}

	// 值类型 调用接收者也是值类型的方法
	fmt.Println(qcrao.howOld())

	// 值类型 调用接收者是指针类型的方法
	qcrao.growUp()
	fmt.Println(qcrao.howOld())

	// ----------------------

	// stefno 是指针类型
	stefno := &Person{age: 100}

	// 指针类型 调用接收者是值类型的方法
	fmt.Println(stefno.howOld())

	// 指针类型 调用接收者也是指针类型的方法
	stefno.growUp()
	fmt.Println(stefno.howOld())

	/*
		总结：
		1、只要传的是结构体指针，即为指针类型调用者，无论接受值是值类型还是指针类型 都修改的是针对指针进行修改 结构体的值也会相应的改变
		2、如果值值传递 值接受 那么方法内修改的结果不会有影响 因为该的都是副本
		3、值传递 指针接受 取地址传入 如果修改 还是针对指针修改 所以原结构体改变
	*/
}

type coder interface {
	code()
	//debug()
}

type Gopher struct {
	language string
}

func (p Gopher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
	p.language = "Python"
}

func (p *Gopher) debug() {
	fmt.Printf("I am debuging %s language\n", p.language)

}
func InterfaceStructure() {
	var c Gopher = Gopher{"Go"}
	c.code()
	c.debug()
	//c.debug()
	// 针对接口 如果传的是值 接受是指针 那么会编译报错 因为接口方法实现的是指针 不是结构体对象

	// 但是反过来 针对接口 如果传的是指针 接受的是值的话 编译会将指针转换成对象 解指针操作 所以不会有问题
	// 针对接口 传指针 接收的是指针 原结构体改变
	// 针对接口 传指针 接受的是值 原结构体不改变 解引用为值
}
