package main

//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	//var has crypto.Hash
//	//hash:=md5.New()
//	//hash.Write([]byte("123456"))
//	//
//	//fmt.Println(hex.EncodeToString(hash.Sum(nil)))
//
//	// data race
//	var nums []int
//
//
//	for i := 0; i < 100; i++ {
//		go func(i int) {
//			nums = append(nums, i)
//		}(i)
//	}
//	//timeout:
//	time.Sleep(time.Second * 2)
//	//fmt.Println(nums)
//	fmt.Println(len(nums))
//}
