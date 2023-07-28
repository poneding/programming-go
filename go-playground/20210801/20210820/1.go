package main

import (
	"fmt"
	"os"
)

func main1() {
	//go func() {
	//	defer fmt.Println("A.defer")	// 可以打印
	//	func(){
	//		defer fmt.Println("B.defer")	// 可以打印
	//		runtime.Goexit()
	//		defer fmt.Println("C.defer")
	//		fmt.Println("B")	// 不能打印
	//	}()
	//	fmt.Println("A")	// 不能打印
	//}()
	//for true {
	//
	//}

	//2
	test()

}

func test() {
	fmt.Println("test.1")
	defer fmt.Println("test.A")
	fmt.Println("test.2")
	os.Exit(1) // 直接退出，不会打印defer
	defer fmt.Println("test.B")
	fmt.Println("test.C")
}
