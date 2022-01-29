package main

import "fmt"

func doubleIt(numToDouble int) int {
	return 2 * numToDouble
}

func squareIt(numToSquare int) int {
	return numToSquare * numToSquare
}

func funFun(functionName func(int) int, variableName int) int {
	return functionName(variableName)
}

func main() {
	fmt.Println("funFun Double:", funFun(doubleIt, 12))
	fmt.Print("funFun Square:", funFun(squareIt, 12))
	fmt.Println("Inline", funFun(func(numToCube int) int { return numToCube * numToCube * numToCube }, 12))
}
