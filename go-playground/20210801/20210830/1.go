package main

import "fmt"

func main() {
	//a := [4]int{1, 2, 3, 4}
	//b := a[0:2]
	//c := a[2:4]
	//b[0] = 10
	//fmt.Println(len(b))
	//fmt.Println(cap(b))
	//c[0] = 30
	//fmt.Println(len(c))
	//fmt.Println(cap(c))
	//fmt.Println(a)
	//b = append(b, 5, 6, 7)
	//fmt.Println(len(b))
	//fmt.Println(cap(b))
	//
	//fmt.Println(len(c))
	//fmt.Println(cap(c))

	//var c chan int
	//var c chan int = make(chan int)
	//var c chan int = make(chan int,1)
	//c <- 1
	//go func() {
	//	<-c
	//}()
	////c <- 1
	//fmt.Println("hello")

	m := make(map[uint]int, 0)
	m[1] = 1
	fmt.Println(m[2])
}
