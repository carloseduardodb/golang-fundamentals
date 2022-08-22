package main

import "fmt"

func closure() func() {
	text := "Dentro da função closure"
	function := func() {
		fmt.Println(text)
	}
	return function
}

func main() {
	text := "Dentro da main"
	fmt.Println(text)

	newFunction := closure()
	newFunction()
}
