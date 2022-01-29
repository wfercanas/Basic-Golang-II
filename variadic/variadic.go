package main

import "fmt"

func varFunc(input ...string) {
	fmt.Println(input)
}

func oneByOne(message string, sliceOfNumbers ...int) int {
	fmt.Println(message)
	sum := 0
	for indexInSlice, sliceElement := range sliceOfNumbers {
		fmt.Print(indexInSlice, sliceElement, "\t")
		sum += sliceElement
	}
	fmt.Println()
	sliceOfNumbers[0] = -1000
	return sum
}

func main() {
	many := []string{"12", "3", "b"}
	varFunc(many...)
	sum := oneByOne("Adding numbers...", 1, 2, 3, 4, 5, -1, 10)
	fmt.Println("Sum:", sum)
	sliceOfNumbers := []int{1, 2, 3}
	sum = oneByOne("Adding numbers...", sliceOfNumbers...)
	fmt.Println(sliceOfNumbers)
}
