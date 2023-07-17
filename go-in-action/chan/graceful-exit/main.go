package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// func main() {
// 	// 运行之后，ctrl+c，立马退出，不会等待task完成
// 	fmt.Println("task start.")
// 	time.Sleep(time.Second * 5) // 假设这是一个需要执行5s的task
// 	fmt.Println("task end.")
// }

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	finishCh := make(chan struct{})

	go func(ctx context.Context, finishCh chan<- struct{}) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stopped.")
				finishCh <- struct{}{}
				return
			default:
				fmt.Println("task start.")
				time.Sleep(time.Second * 5) // 假设这是一个需要执行5s的task
				fmt.Println("task end.")
			}
		}
	}(ctx, finishCh)

	<-sig
	cancel()
	<-finishCh
	fmt.Println("finished.")
}

func RunHttpServer() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Start to handle the request")
		time.Sleep(time.Second * 10)
		fmt.Println("Hello World!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
