package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(1 * time.Second)
	<-timer1.C
	fmt.Println("timer1 fired")

	timer2 := time.NewTimer(1 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 fired")
	}()
	stop := timer2.Stop()
	if stop {
		fmt.Println("timer2 stopped")
	}
	time.Sleep(2 * time.Second)
}

/*
$ go run timers.go
timer1 fired
timer2 stopped
*/
