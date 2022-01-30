package main

import (
	"encoding/json"
	"fmt"
)

func prettyPrint(s interface{}) {
	p, _ := json.MarshalIndent(s, "", "\t")
	fmt.Println(string(p))
}

type Contact struct {
	Name string
	Main Telephone
	Tel  []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func main() {
	contact := new(Contact)
	fmt.Println("Contact:", contact)
	telephone := new(Telephone)
	fmt.Println("Telephone:", telephone)

	if contact.Main == (Telephone{}) {
		fmt.Println("contact.Main is an empty Telephone struct")
	}
	fmt.Println("contact.Main")
	prettyPrint(contact.Main)

	if contact.Tel == nil {
		fmt.Println("contact.Tel is nil.")
	}

	fmt.Println("prettyPrint(contact)")
	prettyPrint(contact)
	fmt.Println("prettyPrint(telephone)")
	prettyPrint(telephone)
}
