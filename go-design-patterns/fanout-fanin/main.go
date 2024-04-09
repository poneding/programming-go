package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	run(3)
}

func worker(id int, jobs <-chan int, result chan<- string) {
	for j := range jobs {
		fmt.Printf("Worker with id %d started job %d\n", id, j)
		time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
		// fan-in
		result <- fmt.Sprintf("Worker with id %d finished job %d", id, j)
	}
}

func run(noOfWorkers int) {
	jobCount := 10
	jobs := make(chan int, jobCount)
	result := make(chan string, jobCount)

	for i := 0; i < noOfWorkers; i++ {
		go worker(i, jobs, result)
	}

	for j := 0; j < jobCount; j++ {
		// fan-out
		jobs <- j
	}
	close(jobs)

	for r := 0; r < jobCount; r++ {
		fmt.Println(<-result)
	}
}
