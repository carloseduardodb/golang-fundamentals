package main

import "fmt"

func functionOne() {
	fmt.Println("FunctionOne")
}

func functionTwo() {
	fmt.Println("FunctionTwo")
}

func main() {
	defer functionOne()
	defer functionTwo()
	fmt.Println("Main")
}
