package recursion

import "fmt"

func Factorial(x int) int {
	if x > 1 {
		result := x * Factorial(x-1)
		return result
	}
	return 1
}

func TryFactorial(x int) {
	fmt.Println(Factorial(x))
}
