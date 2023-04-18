package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// func main(){
// 	run()
// }

func run2() {
	out := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			out <- rand.Intn(100)
		}
		close(out)
	}()

	go func() {
		defer wg.Done()
		for i := range out {
			fmt.Println(i)
		}
	}()
	wg.Wait()
}
