package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name":     "Pedro",
		"lastName": "Silva",
	}

	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"name": {
			"first": "João",
			"last":  "Pedro",
		},
	}

	fmt.Println(user2)
	delete(user2, "name")
	fmt.Println(user2)

	user2["preference"] = map[string]string{
		"name": "cat",
	}

	fmt.Println(user2)
}
