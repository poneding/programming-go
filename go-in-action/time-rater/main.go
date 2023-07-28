package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"sync/atomic"
	"time"
)

var totalQuery int32

func main() {
	go callHandler()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Printf("过去 1 分钟内接口被调用了 %d 次\n", atomic.LoadInt32(&totalQuery)) // 打印统计
		atomic.StoreInt32(&totalQuery, 0)                                  // 重置
	}
}

func handler() {
	atomic.AddInt32(&totalQuery, 1)
	time.Sleep(50 * time.Millisecond)
}

func callHandler() {
	// NewLimiter
	// r：每过多久产生一个令牌
	// b：令牌桶的最大容量
	lim := rate.NewLimiter(rate.Every(100*time.Millisecond), 1)
	for {
		// 1. WaitN 阻塞，直到 lim 允许 n 个事件发生。如果 n 超过了限制器的突发大小，
		// 上下文被取消，或者预期的等待时间超过了上下文的截止日期，则返回错误。
		// 如果速率限制为 Inf，则忽略突发限制。
		// lim.WaitN(context.Background(), 1)
		// handler()

		// 2. ReserveN 返回一个 Reservation 对象表示在 n 个时间发生之前必须等待多长时间
		// reserve := lim.ReserveN(time.Now(), 1)
		// time.Sleep(reserve.Delay())
		// handler()

		// 3. AllowN
		if lim.AllowN(time.Now(), 1) {
			handler()
		}
	}
}
