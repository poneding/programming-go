package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack("hello")
	l.PushBack("world")
	fmt.Println(l)
}
