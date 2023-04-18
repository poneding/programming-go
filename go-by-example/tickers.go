package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	done <- struct{}{}
	fmt.Println("Ticker stopped")
}

/*
$ go run tickers.go
Tick at 2023-02-03 08:28:40.618806 +0800 CST m=+0.501213167
Tick at 2023-02-03 08:28:41.118838 +0800 CST m=+1.001246084
Tick at 2023-02-03 08:28:41.618838 +0800 CST m=+1.501247459
Ticker stopped
*/
