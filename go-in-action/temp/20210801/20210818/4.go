package main

import (
	"fmt"
	"strings"
)

func main() {
	u := User{
		Name: "Jay",
		Age:  40,
	}
	fmt.Printf("%v\n", u)
	fmt.Printf("%+v\n", u)
	fmt.Printf("%#v\n", u)
	fmt.Printf("%T\n", u)

	//str := "hello"
	//r := str[2]
	//fmt.Println(r)
	//str2:=str[3:5]
	//fmt.Println(string(r))
	//fmt.Println(str2)

	// 不建议使用，性能较差
	//var bt bytes.Buffer
	//bt.WriteString("hello")
	//bt.WriteString("world")
	//fmt.Println(bt.String())

	var sb strings.Builder
	sb.WriteString("hello")
	sb.WriteString("world")
	fmt.Println(sb.String())
}

type User struct {
	Name string
	Age  int
}
