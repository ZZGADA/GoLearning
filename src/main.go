package main

import "fmt"

func helloWorldInMain() {
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World after tag v1.0.0")
}

func main() {
	helloWorldInMain()

	fmt.Println("测试 rebase")
	fmt.Println("测试 rebase")
	fmt.Println("测试 rebase")

}
