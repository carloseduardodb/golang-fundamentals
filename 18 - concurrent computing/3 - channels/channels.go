package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go writer("Hello World", channel)
	for message := range channel {
		fmt.Println(message)
	}
}

func writer(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}
	close(channel)
}
