package main

//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//func main() {
//	// 控制goroutine创建的数量
//	m1()
//	m2()
//}
//
//func m1() {
//	// 使用sync.WaitGroup + 有缓冲的channel
//	var wg sync.WaitGroup
//	c := make(chan struct{}, 10)
//	for i := 0; i < 100; i++ {
//		wg.Add(1)
//		c <- struct{}{}
//		go func(i int) {
//			defer wg.Done()
//			time.Sleep(time.Second)
//			fmt.Println("done:", i)
//			<-c
//		}(i)
//	}
//	wg.Wait()
//	fmt.Println("ALL DONE")
//}
//
//func m2() {
//	pool := NewPool(10)
//	for i := 0; i < 100; i++ {
//		pool.Add(1)
//		go func(i int) {
//			defer pool.Done()
//			time.Sleep(time.Second)
//			fmt.Println("done:", i)
//		}(i)
//	}
//	pool.Wait()
//}
//
//type Pool struct {
//	queue chan struct{}
//	wg    *sync.WaitGroup
//}
//
//func NewPool(size int) *Pool {
//	if size <= 0 {
//		size = 1
//	}
//	return &Pool{
//		queue: make(chan struct{}, size),
//		wg:    &sync.WaitGroup{},
//	}
//}
//
//func (pool *Pool) Add(delta int) {
//	if delta > 0 {
//		for i := 0; i < delta; i++ {
//			pool.queue <- struct{}{}
//		}
//	} else {
//		for i := 0; i < delta; i++ {
//			<-pool.queue
//		}
//	}
//	pool.wg.Add(delta)
//}
//
//func (pool *Pool) Done() {
//	<-pool.queue
//	pool.wg.Done()
//}
//
//func (pool *Pool) Wait() {
//	pool.wg.Wait()
//}
