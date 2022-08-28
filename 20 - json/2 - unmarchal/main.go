package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json:"-"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
}

func main() {
	dogInJSON := `{"name":"Fido","breed":"Labrador","age":3}`
	var d dog
	if error := json.Unmarshal([]byte(dogInJSON), &d); error != nil {
		log.Fatal(error)
	}
	fmt.Println(d)
	dog2InJSON := `{"name":"Fido","breed":"Labrador"}`
	d2 := make(map[string]string)
	if error := json.Unmarshal([]byte(dog2InJSON), &d2); error != nil {
		log.Fatal(error)
	}
	fmt.Println(d2)
}
