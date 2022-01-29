package main

import "fmt"

func doubleSquare(firstNum int) (int, int) {
	return firstNum * 2, firstNum * firstNum
}

func namedMinMax(firstNum, secondNum int) (min, max int) {
	if firstNum > secondNum {
		min = secondNum
		max = firstNum
	} else {
		min = firstNum
		max = secondNum
	}

	return
}

func minMax(firstNum, secondNum int) (min, max int) {
	if firstNum > secondNum {
		min = secondNum
		max = firstNum
	} else {
		min = firstNum
		max = secondNum
	}

	return min, max
}

func main() {
	secondNum := 10

	square := func(numberToSquare int) int {
		return numberToSquare * numberToSquare
	}
	fmt.Println("The square of", secondNum, "is", square(secondNum))

	double := func(numberToDouble int) int {
		return 2 * numberToDouble
	}
	fmt.Println("The double of", secondNum, "is", double(secondNum))

	fmt.Println(doubleSquare(secondNum))
	doubledNumber, squaredNumber := doubleSquare(secondNum)
	fmt.Println(doubledNumber, squaredNumber)

	value1 := -10
	value2 := -1
	fmt.Println(minMax(value1, value2))
	min, max := minMax(value1, value2)
	fmt.Println(min, max)
	fmt.Println(namedMinMax(value1, value2))
	min, max = namedMinMax(value1, value2)
	fmt.Println(min, max)
}
