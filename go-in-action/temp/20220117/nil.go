package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 空切片和nil切片
	// 空切片的引用地址指向zero数组，所有空切片的引用地址都指向同一个zero数组，所以所有空切片都可以判等
	// nil切片没有引用地址，直接使用会抛异常 但是len()函数和cap函数是可以正常使用的，值均为0

	var s1 []int
	s2 := make([]int, 0)
	s4 := make([]int, 0)
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
	fmt.Println(len(s4), cap(s4))

	fmt.Printf("s1 pointer:%+v, s2 pointer:%+v, s4 pointer:%+v, \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)), *(*reflect.SliceHeader)(unsafe.Pointer(&s2)), *(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s1))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data)
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s4))).Data)
}
