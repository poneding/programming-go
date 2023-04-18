package main

import "fmt"

func main() {
	jobs := make(chan int, 50)
	result := make(chan int, 50)
	go worker(jobs, result)
	for i := 0; i < 50; i++ {
		jobs <- i
	}
	close(jobs)
	for r := range result {
		fmt.Println(r)
	}
}

func worker(jobs <-chan int, result chan<- int) {
	for n := range jobs {
		result <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
