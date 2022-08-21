package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")
	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando J", j)
	// 	time.Sleep(time.Second)
	// }

	names := []string{"João", "Davi", "Lucas"}
	for _, name := range names {
		fmt.Println(name)
	}

	for _, letter := range "Teste" {
		fmt.Println(string(letter))
	}

	user := map[string]string{
		"name":      "Leonardo",
		"sobrenome": "Silva",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}
}
