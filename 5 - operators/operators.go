package main

import "fmt"

func main() {
	// ARITMETICOS
	var sum int16 = 1 + 2
	var sub int16 = 1 - 2
	var division int16 = 10 / 4
	var multiplication int16 = 10 * 5
	var restOfDivision int16 = 10 % 2

	fmt.Println(
		sum,
		sub,
		division,
		multiplication,
		restOfDivision,
	)

	/*
		se eu tenho 2 números que somados estouram o tamanho do inteiro
		na memoria eu recebo negativo se eu coloco esses números como float
		a aplicação consegue fazer a conta e retornar um valor
	*/
	var test1 int64 = 999999999999999999
	var test2 int64 = 999999999999999999

	sum1 := test1 * test2

	fmt.Println("Multiplicação: ", sum1)

	// FINISHED ARITHMETIC

	// ATRIBUIÇÕES
	var variable1 string = "String"
	variable2 := "String2"
	fmt.Println(variable1, variable2)
	// FIM OPERADORES DE ATRIBUIÇÃO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// FIM OPERADORES RELACIONAIS

	// OPERADORES LÓGICOS
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	// FIM DOS OPERADORES LÓGICOS

	// OPERADORES UNÁRIOS
	number := 10
	number++
	number *= 3
	number /= 4
	number -= 2
	number %= 3
	fmt.Println(number)
	// FIM DOS OPERADORES UNÁRIOS
}
