package main

import "fmt"

func main() {
	fmt.Println("control structures")

	number := 10

	if number > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if myNumber := number; myNumber > 0 {
		fmt.Println("Número é maior que zero")
	}
}
