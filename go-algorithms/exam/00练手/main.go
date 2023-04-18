package main

import "fmt"

func main() {
	c := make(chan int, 1)
	go func() {
		c <- 1
		close(c)
	}()

	for i := 0; i < 2; i++ {
		v := <-c
		fmt.Println(v)
	}
}
