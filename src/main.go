package main

import lambdafunction "basicLearning/src/lambdaFunction"

func usingLambda() {
	for i := 0; i < 10; i++ {
		lambdafunction.UsingLambda()
	}

}

func main() {
	//fmt.Println("Hello World")
	//arrayAndSlice.ArrayLearning()
	usingLambda()
}
