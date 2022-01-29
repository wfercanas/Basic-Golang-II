package main

import "fmt"

func main() {
	for loopIndex := 0; loopIndex < 20; loopIndex++ {
		if loopIndex%10 == 0 {
			// Prevents this specific round from going on and skips to the next loopIndex
			continue
		}

		if loopIndex == 19 {
			// Breaks the whole for-loop
			break
		}
		fmt.Print(loopIndex, " ")
	}
	fmt.Println()

	fmt.Println("Break with DESC:")
	// Use break to exit the for loop
	loopIndex := 10
	for {
		if loopIndex < 0 {
			break
		}
		fmt.Print(loopIndex, " ")
		loopIndex--
	}
	fmt.Println()

	fmt.Println("While-like loop:")
	// This is similar to a while(true) do something loop
	loopIndex = 0
	anExpression := true

	for ok := true; ok; ok = anExpression {
		if loopIndex > 10 {
			anExpression = false
		}

		fmt.Print(loopIndex, " ")
		loopIndex++
	}
	fmt.Println()

	anArray := [5]int{0, 1, -1, 2, -2}
	for loopIndex, value := range anArray {
		fmt.Println("index:", loopIndex, "value: ", value)
	}
}
