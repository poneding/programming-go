package main

import (
	"fmt"
	"time"
)

func main() {
	// execTime := time.Second * 4
	execTime := time.Second
	timeout := time.Second * 2

	ch := make(chan struct{}, 1)
	go func() {
		time.Sleep(execTime)
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		fmt.Println("exec done.")
	case <-time.After(timeout):
		fmt.Println("exec timeout.")
	}
}
