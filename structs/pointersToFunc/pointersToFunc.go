package main

import "fmt"

type Employee struct {
	FirstName  string
	employeeID int
}

func ChangeEmployeeID(e *Employee, newID int) {
	e.employeeID = newID
}

func main() {
	employeePointer1 := &Employee{"Nathan", 8124011}
	fmt.Println("Pointer1:", *employeePointer1) //  {Nathan 8124011}
	ChangeEmployeeID(employeePointer1, 1012843)
	fmt.Println("Pointer1:", *employeePointer1) // {Nathan 1012843}
}
