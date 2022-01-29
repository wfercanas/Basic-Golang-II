package main

import "fmt"

type Employee struct {
	FirstName  string
	employeeID int
}

func (e Employee) PrintGreeting() {
	fmt.Println("My name is", e.FirstName, "and my employee ID is", e.employeeID)
}

func main() {
	employee1 := Employee{"Nathan", 8124011}
	employee1.PrintGreeting()
}
