package main

import (
	"fmt"
	"sync"
)

type LinkNode struct {
	v    int
	next *LinkNode
}

type LinkStack struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

func (stack *LinkStack) Put(v int) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.root == nil {
		stack.root = &LinkNode{v: v}
	} else {
		new := &LinkNode{v: v}
		new.next = stack.root
		stack.root = new
	}
	stack.size++
}

func (stack *LinkStack) Pop() int {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size == 0 {
		panic("empty panic")
	}

	v := stack.root.v
	stack.root = stack.root.next
	stack.size--
	return v
}

func (stack *LinkStack) Peek() int {
	if stack.size == 0 {
		panic("empty panic")
	}
	return stack.root.v
}

func (stack *LinkStack) Size() int {
	return stack.size
}

func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}

func main() {
	linkStack := new(LinkStack)
	linkStack.Put(1001)
	linkStack.Put(1002)
	linkStack.Put(1003)
	fmt.Println("size:", linkStack.Size())
	fmt.Println("peek:", linkStack.Peek())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("peek:", linkStack.Peek())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("size:", linkStack.Size())
	linkStack.Put(1005)
	fmt.Println("pop:", linkStack.Pop())
}
