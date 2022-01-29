package main

import "fmt"

type Employee struct {
	FirstName  string
	employeeID int
}

func ChangeEmployeeID(e Employee, newID int) {
	e.employeeID = newID
}

func main() {
	var nathan Employee
	fmt.Println(nathan)
	nathan = Employee{FirstName: "Nathan", employeeID: 8124011}
	fmt.Println(nathan)

	var heather Employee = Employee{FirstName: "Heather"}
	fmt.Println(heather)

	mihalis := Employee{"Mihalis", 1910234}
	fmt.Println(mihalis)

	// Comparison of structs
	employee2 := Employee{"Mihalis", 1910234}
	fmt.Println(mihalis == employee2) // true

	// Accessing a field
	fmt.Println("my name is", mihalis.FirstName, "and my employee ID is", mihalis.employeeID)

	// To assign is to Copy
	andy := nathan
	fmt.Println("Andy:", andy)
	andy.FirstName = "Andy"
	andy.employeeID = 122235
	nathan.FirstName = "Nate"
	fmt.Println("Andy", andy)
	fmt.Println("Nathan", nathan)

	ChangeEmployeeID(andy, 999876)
	fmt.Println("Andy", andy)

}
