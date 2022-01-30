// A Struct used for reading a text file that contains data in the JSON format and creating data in the JSON format

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func loadFromJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}

	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
}

func saveToJSON(filename *os.File, key interface{}) {
	encodeJSON := json.NewEncoder(filename)
	err := encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename!")
		return
	}

	filename := arguments[1]
	var myRecord Record
	err := loadFromJSON(filename, &myRecord)
	fmt.Println("JSON file loaded into struct:")
	if err == nil {
		fmt.Println(myRecord)
	} else {
		fmt.Println(err)
	}

	myRecord = Record{
		Name:    "Mihalis",
		Surname: "Tsoukalos",
		Tel:     []Telephone{{Mobile: true, Number: "1234-5687"}, {Mobile: true, Number: "6789-abcd"}, {Mobile: false, Number: "FAVA-5689"}},
	}

	fmt.Println("struct saved to JSON:")
	saveToJSON(os.Stdout, myRecord)
}
