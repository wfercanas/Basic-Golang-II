package main

import "fmt"

type Employee struct {
	FirstName  string
	employeeID int
}

func (e *Employee) ChangeEmployeeID(newID int) {
	e.employeeID = newID
}

func main() {
	// With a pointer
	var employeePointer1 *Employee = &Employee{"Nathan", 8124011}
	fmt.Println(*employeePointer1) // {Nathan 8124011}
	employeePointer1.ChangeEmployeeID(1017193)
	fmt.Println(*employeePointer1) // {Nathan 1017193}

	// Without a pointer
	var employee2 Employee = Employee{"Nathan", 8124011}
	fmt.Println(employee2)
	employee2.ChangeEmployeeID(1017193)
	fmt.Println(employee2)
}
