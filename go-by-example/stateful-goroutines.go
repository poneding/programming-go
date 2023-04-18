package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

func main() {

	var readOps, writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for i := 0; i < 100; i++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read // 发送写消息
				<-read.resp   // 写结果
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Printf("readOpsFinal: %v\n", readOpsFinal)

	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Printf("writeOpsFinal: %v\n", writeOpsFinal)
}

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

/*
$  go run stateful-goroutines.go
readOpsFinal: 86254
writeOpsFinal: 8662
*/
