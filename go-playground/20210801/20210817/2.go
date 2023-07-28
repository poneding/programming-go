package main

//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//func main() {
//	c1, c2 := make(chan int), make(chan int)
//	longTimeRequest(c1)
//	longTimeRequest(c2)
//
//	fmt.Println(sumSquares(<-c1, <-c2))
//}
//
//func longTimeRequest(c chan<- int) {
//	go func() {
//		time.Sleep(time.Second * 5)
//		c <- rand.Int()
//	}()
//}
//
//func sumSquares(a, b int) int {
//	return a*a + b*b
//}
