package main

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"
)

type ServerConn struct {
	Connection net.Conn
	ID         string
	Open       bool
}

func ShowConnection(p *atomic.Pointer[ServerConn]) {
	for {
		time.Sleep(10 * time.Second)
		fmt.Println(p, p.Load())
	}

}
func main() {
	c := make(chan bool)
	p := atomic.Pointer[ServerConn]{}
	s := ServerConn{ID: "first_conn"}
	p.Store(&s)
	go ShowConnection(&p)
	go func() {
		for {
			time.Sleep(13 * time.Second)
			newConn := ServerConn{ID: "new_conn"}
			p.Swap(&newConn)
		}
	}()
	<-c
}

/*
$ go run atomic-pointer/atomic-pointer.go
&{[] {} 0x1400010c420} &{<nil> first_conn false}
&{[] {} 0x140000a8030} &{<nil> new_conn false}
&{[] {} 0x1400006e030} &{<nil> new_conn false}
&{[] {} 0x140000a8060} &{<nil> new_conn false}
*/
