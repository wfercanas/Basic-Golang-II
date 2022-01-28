package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// fmt.Println("Args:", os.Args)              // [./cla arg1 arg2]
	// fmt.Println("len(os.Args):", len(os.Args)) // 3

	if len(os.Args) == 1 {
		// Checks if ./cla was executed with no extra parameters, if so:
		fmt.Println("Please give one or more integers")
		return
	}

	var min, max int

	arguments := os.Args
	// fmt.Println("Arguments:", arguments)
	// fmt.Printf("Arguments is of type %T \n", arguments)

	temp, err := strconv.Atoi(arguments[1])
	// fmt.Println("temp:", temp)
	// fmt.Println("err:", err)

	if err != nil {
		fmt.Println("Error encountered, exiting:")
		fmt.Println(err)
		return
	} else {
		min = temp
		max = temp
	}

	for i := 2; i < len(arguments); i++ {
		n, err := strconv.Atoi(arguments[i])
		if err != nil {
			fmt.Println("Error encountered, exiting:")
			fmt.Println(err)
			return
		}

		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
