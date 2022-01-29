package main

import "fmt"

type Employee struct {
	FirstName  string
	employeeID int
}

func main() {
	var employeePointer1 *Employee = &Employee{"Nathan", 1201921}
	fmt.Println("Getting a specific struct field:", (*employeePointer1).FirstName) // Nathan
	fmt.Println("With implicit dereferencing:", employeePointer1.FirstName)        // Nathan

	employeePointer2 := employeePointer1
	employeePointer2.FirstName = "Nate"
	fmt.Println("FirstName for employeePointer2:", employeePointer2.FirstName) // Nate
	fmt.Println("FirstName for employeePointer1:", employeePointer1.FirstName) // Nate
}
