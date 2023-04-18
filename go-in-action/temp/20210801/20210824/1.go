package main

import (
	"fmt"
	"math/rand"
)

func main1() {
	// 随机返回0和1
	// 第一种使用rand.Intn(2)
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(2))
	}
	fmt.Println("------------------")

	// 第二种使用select随机返回
	for i := 0; i < 10; i++ {
		fmt.Println(random())
	}
}

func random() int {
	ch := make(chan int)
	go func() {
		select {
		case ch <- 0:
		case ch <- 1:
		}
	}()

	return <-ch
}
