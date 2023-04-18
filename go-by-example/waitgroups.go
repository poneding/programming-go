package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// defer wg.Wait()

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			worker(i)
		}(i)
	}

	wg.Wait()
}

func worker(id int) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d done\n", id)
}

/*
$ go run waitgroups.go
worker 5 starting
worker 4 starting
worker 1 starting
worker 2 starting
worker 3 starting
worker 3 done
worker 5 done
worker 4 done
worker 2 done
worker 1 done
*/
