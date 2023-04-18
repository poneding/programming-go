package main

import (
	"fmt"
	"math/rand"
	"time"
)

// rand7()是一个可以随机返回1~7的函数，基于此，
// 实现一个随机返回1~10的函数
func main() {
	fmt.Println(rand10())
	fmt.Println(rand10())
	fmt.Println(rand10())
	fmt.Println(rand10())
	fmt.Println(rand10())
	fmt.Println(rand10())
	fmt.Println(rand10())
	fmt.Println(rand10())
	fmt.Println(rand10())
}

func rand7() int {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(1 * time.Millisecond)
	return rand.Intn(7) + 1
}

func rand10() int {
	a := rand7()
	b := rand7()
	num := a + (b-1)*7
	for num > 40 {
		a = rand7()
		b = rand7()
		num = a + (b-1)*7
	}
	return num % 10
}
