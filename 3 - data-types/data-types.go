package main

import (
	"errors"
	"fmt"
)

func main() {
	var number int64 = -1000000000000000000
	fmt.Println(number)

	var number2 uint64 = 10000
	fmt.Println(number2)

	// alias
	// INT32 = RUNE
	var number3 rune = 123456
	fmt.Println(number3)

	// BYTE = UINT8
	var number4 byte = 123
	fmt.Println(number4)

	var number5 float32 = 0.1 * 0.2
	fmt.Println(number5)

	// FINISHED NUMBERS

	var str string = "Text"
	fmt.Println(str)

	str2 := "Text2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	// FINISHED STRING

	var text int64
	fmt.Println(text)

	var boolean1 bool
	fmt.Println(boolean1)

	var error1 error = errors.New("Internal Server Error")
	fmt.Println(error1)
}
