package main

import "fmt"

func test() string {
	return "Hello"
}

func calcMath(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	fmt.Println(test())

	var f = func(param, param2 string) {
		fmt.Println(param)
	}

	f("Hello2", "Hello2")

	resultSum, _ := calcMath(10, 15)
	fmt.Println(resultSum)
}
