package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic(1)
	generic("Hello")
	generic(true)

	mapping := map[interface{}]interface{}{
		1:            "um",
		2:            "dois",
		"String":     "String",
		float32(3.0): true,
		true:         float64(3.0),
	}

	fmt.Println(mapping)
}
