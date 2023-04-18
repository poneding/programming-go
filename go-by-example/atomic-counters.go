package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				atomic.AddUint64(&ops, 1)
			}
		}()
	}
	wg.Wait()

	fmt.Println("ops:", ops)
}

/*
$ go run atomic-counters.go
ops: 5000
*/
