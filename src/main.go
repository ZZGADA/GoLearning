package main

import (
	lambdafunction "basicLearning/src/lambdaFunction"
	"fmt"
)

func usingLambda() {
	for i := 0; i < 10; i++ {
		lambdafunction.UsingLambda()
	}

}

func main() {
	//fmt.Println("Hello World")
	//arrayAndSlice.ArrayLearning()
	//usingLambda()
	//maplearning.TryMap()
	//recursion.TryFactorial(6)
	//errorcatch.TryDivideError(0)
	//basicgrammer.BasicGrammar()
	//structure.TryUsingInterface()
	//structure.UsingStructureFunc()
	//structure.UsingJsonSerialize()
	//structure.StructureValueAndPointer()
	//structure.InterfaceStructure()
	//errorcatch.TrySelfDefineError()

	remember := map[int]int{}
	remember[10] = 1
	fmt.Println(remember[10])
	fmt.Println(remember[11])
}
