package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//func init() {
//	rand.Seed(time.Now().Unix())
//}
//func main() {
//	c1, c2 := longTimeRequest(), longTimeRequest()
//	fmt.Println(sumSquares(<-c1, <-c2))
//}
//
//// 返回一个单向接收channel
//func longTimeRequest() <-chan int {
//	r := make(chan int)
//	go func() {
//		time.Sleep(time.Second * 5)
//		r <- rand.Int()
//	}()
//	return r
//}
//
//func sumSquares(a, b int) int {
//	return a*a + b*b
//}
//
//// 并发同步：
//// 控制若干并发计算：避免他们之间产生数据竞争的现象；避免他们无所事事的时候消耗资源
