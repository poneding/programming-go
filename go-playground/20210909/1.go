package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a1 := [2]string{"Hello", "World"}
	fmt.Printf("数组指针：%p\n", &a1)
	ap(a1)

	s1 := a1[0:1]
	fmt.Printf("切片指针：%v\n",(*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)
	sp(s1)
}

func ap(a [2]string) {
	fmt.Printf("数组指针：%p\n", &a)
}

func sp(s []string) {
	fmt.Printf("切片指针：%v\n",(*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)
}
