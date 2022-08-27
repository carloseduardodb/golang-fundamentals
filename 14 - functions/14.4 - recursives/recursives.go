package main

import "fmt"

// fibonacci - Função recursiva que calcula o n-ésimo número da sequência de Fibonacci.
func fibonacci(position uint) uint {
	if position < 2 {
		return position
	}
	return fibonacci(position-1) + fibonacci(position-2)
}

func main() {
	fmt.Println("Funções recursivas")
	const position uint = 40
	for i := uint(1); i < position; i++ {
		fmt.Println(fibonacci(i))
	}
}
