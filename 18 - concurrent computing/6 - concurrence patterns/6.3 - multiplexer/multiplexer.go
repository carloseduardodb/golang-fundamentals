package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplexer(writer("Hello World!"), writer("Hello World2!"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexer(channelToInput1, channelToInput2 <-chan string) <-chan string {
	channelToOutput := make(chan string)
	go func() {
		for {
			select {
			case value := <-channelToInput1:
				channelToOutput <- value
			case value := <-channelToInput2:
				channelToOutput <- value
			}
		}
	}()
	return channelToOutput
}

func writer(text string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- fmt.Sprintf("Receive value: %s", text)
			time.Sleep(time.Microsecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return channel
}
