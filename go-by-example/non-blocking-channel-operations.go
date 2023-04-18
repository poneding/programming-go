package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("receieved message", msg)
	default:
		fmt.Println("no message receieved")
	}

	msg := "hello"
	select {
	// 不能送到 messages chan 中，因为 messages 是一个无缓冲 chan，且没有接收方
	case messages <- msg:
		fmt.Println("sent message", msg)
	// 所以这里会走到 default case
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("receieved message", msg)
	case sig := <-signals:
		fmt.Println("received singal", sig)
	default:
		fmt.Println("no activity")
	}
}

/*
$ go run non-blocking-channel-operations.go
no message receieved
no message sent
no activity
*/
