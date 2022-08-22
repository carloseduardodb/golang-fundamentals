package main

import "fmt"

func main() {
	fmt.Println("Funções recursivas")
	const position uint = 10
	fmt.Printf("O %dº número da sequência de Fibonacci é: %d\n", position, fibonacci(position))
}

// fibonacci - Função recursiva que calcula o n-ésimo número da sequência de Fibonacci.
func fibonacci(position uint) uint {
	if position < 2 {
		return position
	}
	return fibonacci(position-1) + fibonacci(position-2)
}
