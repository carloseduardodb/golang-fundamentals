package main

import "fmt"

func main() {
	tasks := make(chan int, 60)
	results := make(chan int, 60)
	go worker(tasks, results)
	go worker(tasks, results)
	for i := 0; i < 60; i++ {
		tasks <- i
	}
	close(tasks)
	for i := 0; i < 60; i++ {
		fmt.Println(<-results)
	}
}

func worker(tasks <-chan int, results chan<- int) {
	for number := range tasks {
		results <- fibonacci(number)
	}
}

func fibonacci(position int) int {
	if position < 2 {
		return position
	}
	return fibonacci(position-1) + fibonacci(position-2)
}
