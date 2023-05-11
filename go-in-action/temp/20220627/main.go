package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan struct{})
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("gogogo")
		}
	}()
	<-stop
}
