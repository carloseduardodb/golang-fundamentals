package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var variable1 int = 10
	var variable2 int = variable1

	fmt.Println(variable1, variable2)

	variable2++
	fmt.Println(variable1, variable2)

	// PONTEIRO É UM REFERENCIA DE MEMORIA
	var variable3 int
	var pointer *int // cria um ponteiro

	variable3 = 100

	pointer = &variable3 // atribui ao ponteiro a referencia de memoria da variável 3

	fmt.Println(variable3, pointer)

	/** quando a variável 3 muda de valor, não quer dizer que ela mudou
	  * o seu local na memoria e sim que ela trocou o seu valor
	  * como o ponteiro está apontado para endereço de memoria da variável 3
		* quando eu printo o ponteiro eu printo o valor atual da variavel 3
	*/
	variable3 = 150

	fmt.Println(variable3, *pointer) // desreferenciação

}
