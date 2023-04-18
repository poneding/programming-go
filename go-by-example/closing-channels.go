package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan struct{})

	go func() {
		for {
			if job, ok := <-jobs; ok {
				fmt.Println("received job", job)
			} else {
				// jobs chan 关闭，ok 为 false，job 为类型的零值
				fmt.Println("reveived all jobs")
				done <- struct{}{}
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}

/*
$ go run closing-channels.go
sent job 1
sent job 2
sent job 3
sent all jobs
received job 1
received job 2
received job 3
reveived all jobs
*/
