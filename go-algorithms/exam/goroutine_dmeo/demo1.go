package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg       sync.WaitGroup
	execTime time.Duration = time.Second
)

func finishReq(timeout time.Duration) int {
	// 如果是无缓冲channel，会产生goroutine死锁
	// 因为超时时间比子goroutine执行之间短，select走了第二个case，此时
	// channel没有接收者，而channel又是没有缓冲的，所以子goroutine阻塞
	// 这里主goroutine又使用了wg.wait等待子goroutine，也被阻塞，所以死锁
	//ch := make(chan int)

	// 解决方法可以使用一个有缓冲的channel
	ch := make(chan int, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(execTime)
		ch <- 1
	}()

	select {
	case res := <-ch:
		return res
	case <-time.After(timeout):
		return -1
	}

}

func main1() {
	timeout := time.Millisecond * 100
	res := finishReq(timeout)
	fmt.Println(res)
	wg.Wait()
}
