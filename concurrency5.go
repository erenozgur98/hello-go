package main

import (
	"fmt"
)

// worker pull
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// if we add more workers, the work will be done faster but the CPU usage goes up 300%
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}

	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

// we'll only receive from jobs channel and we only send from results channel to reduce chance of having bugs
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n - 1) + fib(n - 2)
}