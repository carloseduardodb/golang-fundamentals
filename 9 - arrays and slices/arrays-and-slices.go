package main

import (
	"fmt"
)

func main() {
	fmt.Println("Array and Slices")

	var array1 [5]string
	array1[0] = "Position 1"
	fmt.Println(array1)

	array2 := [5]string{"Position 1", "Position 2", "Position 3", "Position 4", "Position 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// slice
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)
}
