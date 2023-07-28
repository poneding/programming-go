package main

//import (
//	"context"
//	"fmt"
//)
//
//func main() {
//	// 通过context.WithCancel来避免goroutine内存泄漏
//	ctx, cancel := context.WithCancel(context.Background())
//	ch := make(chan int)
//	go func(ctx context.Context) {
//		for i := 0; ; i++ {
//			select {
//			case <-ctx.Done():
//				return
//			case ch <- i:
//			}
//		}
//	}(ctx)
//
//	for v := range ch {
//		fmt.Println(v)
//		if v == 5 {
//			cancel()
//			break
//		}
//	}
//}
