package main

import "fmt"

func getPointer(varToPointer *float64) float64 {
	return *varToPointer * *varToPointer
}

func returnPointer(testValue int) *int {
	squareTheBestValue := testValue * testValue
	return &squareTheBestValue
}

func main() {
	testValue := -12.12
	fmt.Println(getPointer(&testValue))
	testValue = -12
	fmt.Println(getPointer(&testValue))

	theSquare := returnPointer(10)
	fmt.Println("sq value:", *theSquare)
	fmt.Println("sq memory address:", theSquare)
}
