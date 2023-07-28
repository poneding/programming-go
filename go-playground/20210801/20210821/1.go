package main

import (
	"sync"
	"time"
)

func main() {

}

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	ch := make(chan struct{})
	go func() {
		wg.Wait()
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		return true
	case <-time.After(timeout):
		return false
	}
}
