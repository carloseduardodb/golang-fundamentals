package main

import (
	"fmt"
	"time"
)

func main() {
	go writer("Hello World") // goroutine
	writer("Go Programming")
	time.Sleep(time.Second * 5)
}

func writer(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
