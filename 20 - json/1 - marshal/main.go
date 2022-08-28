package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
}

func main() {
	d := dog{
		Name:  "Fido",
		Breed: "Labrador",
		Age:   2,
	}
	fmt.Println(d)
	dogJSON, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dogJSON))
}
