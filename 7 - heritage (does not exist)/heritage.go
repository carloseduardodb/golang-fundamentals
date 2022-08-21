package main

import "fmt"

type People struct {
	name     string
	lastName string
	age      uint8
	height   uint8
}

type Student struct {
	People
	course  string
	college string
}

func main() {
	fmt.Println("Hello heritage")

	p1 := People{"João", "Pedro", 20, 178}
	fmt.Println(p1)

	s1 := Student{p1, "Engenharia", "USP"}
	fmt.Println(s1)
}
