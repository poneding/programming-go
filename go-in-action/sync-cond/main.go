package main

import (
	"log"
	"sync"
	"time"
)

var done bool

func read(name string, cond *sync.Cond) {
	cond.L.Lock()
	defer cond.L.Unlock()

	for !done {
		cond.Wait()
	}

	log.Println("reading", name)
}

func write(name string, cond *sync.Cond) {
	log.Println("writing", name)
	<-time.Tick(time.Second)

	cond.L.Lock()
	defer cond.L.Unlock()

	done = true
	// cond.Signal()
	cond.Broadcast()
}

// var c chan struct{}

// func read2(name string) {
// 	<-c
// 	log.Println("readding", name)
// }

// func write2(name string) {
// 	log.Println("writing", name)
//  <-time.Tick(time.Second)
// 	c <- struct{}{}
// }

func main() {

	// writer 发布了，reader 才能读取

	// 使用 sync.Cond 实现
	cond := sync.NewCond(&sync.Mutex{})
	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	go read("reader4", cond)
	go read("reader5", cond)
	write("writer1", cond)

	// channel 通信只能是一对一的，不能是一对多的，所以这种方式行不通
	// go read2("reader1")
	// // go read2("reader2")
	// write2("writer1")

	<-time.Tick(time.Second * 5)
}
