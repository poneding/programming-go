package main

import "fmt"

// any 类型转换成具体的类型
func main() {
	var a = []any{10.123, "234"}

	// var b float64
	var b = a[0].(float64)

	fmt.Println(b)
}
