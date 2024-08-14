package main

import (
	lambdafunction "basicLearning/src/lambdafunction"
	"fmt"
)

func usingLambda() {
	for i := 0; i < 10; i++ {
		lambdafunction.UsingLambda()
	}

}

func main() {
	fmt.Println("Hello World")
}
