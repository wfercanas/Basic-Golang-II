package main

import "fmt"

func squareFunction() func() int {
	numToSquare := 0
	return func() int {
		numToSquare++
		return numToSquare * numToSquare
	}
}

func main() {
	square1 := squareFunction()
	square2 := squareFunction()

	fmt.Println("First Call to square1:", square1())
	fmt.Println("Second Call to square1:", square1())
	fmt.Println("First Call to square2:", square2())
	fmt.Println("Third Call to square1:", square1())
}
