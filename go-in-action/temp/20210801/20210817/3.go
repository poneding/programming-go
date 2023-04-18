package main

//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//func init() {
//	rand.Seed(time.Now().Unix())
//}
//func source(c chan<- int) {
//	a, b := rand.Int(), rand.Int()
//
//	time.Sleep(time.Duration(b) * 3)
//
//	c <- a
//}
//
//func main() {
//	startTime := time.Now()
//	c := make(chan int, 5)
//	for i := 0; i < cap(c); i++ {
//		go source(c)
//	}
//
//	// 最快回应，只有第一个写入channel的值会被接收到
//	rnd := <-c
//	fmt.Println(time.Since(startTime))
//	fmt.Println(rnd)
//}
