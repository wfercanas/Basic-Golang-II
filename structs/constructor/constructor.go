// Another way for creating structs is:
package main

import "fmt"

type Employee struct {
	FirstName  string
	employeeID int
}

func NewEmployee(name string, employeeID int) *Employee {
	if employeeID <= 0 {
		return nil
	}
	return &Employee{name, employeeID}
}

func main() {
	employeePointer1 := NewEmployee("Nathan", 8124011)
	fmt.Println(*employeePointer1) // {Nathan 8124011}
}
