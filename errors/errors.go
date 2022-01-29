package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	customError := errors.New("My Custom Error!")
	if customError.Error() == "My Custom Error!" {
		fmt.Println("!!")
	}

	stringToConvert1 := "123"
	stringToConvert2 := "43W"
	_, err := strconv.Atoi(stringToConvert1)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = strconv.Atoi(stringToConvert2)
	if err != nil {
		fmt.Println(err)
		return
	}
}
