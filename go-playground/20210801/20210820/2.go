package main

//import (
//	"fmt"
//	"strings"
//	"sync"
//)
//
//func main() {
//	fmt.Println(strings.Fields("Hello World")) // [Hello World]
//	fmt.Println(strings.FieldsFunc("Hello World", func(r rune) bool {
//		return r == 'o'
//	})) // [Hell  W rld]
//
//	m := sync.Map{}
//	m.Store("Hello", "World")
//	m.Store("Num", 123)
//	m.Store("Assert", true)
//
//	fmt.Println(m.Load("Num"))
//	fmt.Println(m.Load("Hello"))
//	m.Range(func(key, value interface{}) bool {
//		fmt.Println(key, value)
//		return true
//	})
//}
