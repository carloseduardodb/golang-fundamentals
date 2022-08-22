package main

import "fmt"

func main() {
	returnFunction := func(text string) string {
		return fmt.Sprintf("Received: %s", text)
	}("Passing params")

	fmt.Println(returnFunction)
}
