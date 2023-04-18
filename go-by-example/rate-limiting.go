package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.NewTicker(200 * time.Millisecond)

	for req := range requests {
		<-limiter.C
		fmt.Println("request", req, time.Now())
	}

	fmt.Println()

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

/*
$ go run rate-limiting.go
request 1 2023-02-03 09:09:08.460961 +0800 CST m=+0.201213501
request 2 2023-02-03 09:09:08.66103 +0800 CST m=+0.401282501
request 3 2023-02-03 09:09:08.860849 +0800 CST m=+0.601100959
request 4 2023-02-03 09:09:09.060964 +0800 CST m=+0.801216917
request 5 2023-02-03 09:09:09.26104 +0800 CST m=+1.001292501

request 1 2023-02-03 09:09:09.261324 +0800 CST m=+1.001576709
request 2 2023-02-03 09:09:09.261337 +0800 CST m=+1.001590292
request 3 2023-02-03 09:09:09.261344 +0800 CST m=+1.001597001
request 4 2023-02-03 09:09:09.461459 +0800 CST m=+1.201711667
request 5 2023-02-03 09:09:09.661489 +0800 CST m=+1.401742251
*/
