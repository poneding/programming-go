package main

import (
	"fmt"
	"sync"
)

type ArrayStack struct {
	data []int
	size int
	lock sync.Mutex
}

func (stack *ArrayStack) Put(v int) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.data = append(stack.data, v)
	stack.size++
}

func (stack *ArrayStack) Pop() int {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size == 0 {
		panic("empty stack")
	}

	v := stack.data[stack.size-1]
	newData := make([]int, stack.size-1, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newData[i] = stack.data[i]
	}
	stack.data = newData
	stack.size--
	return v
}

func (stack *ArrayStack) Peek() int {
	if stack.size == 0 {
		panic("empty stack")
	}

	return stack.data[stack.size-1]
}

func (stack *ArrayStack) Size() int {
	return stack.size
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}

func main1() {
	arrayStack := new(ArrayStack)
	arrayStack.Put(1001)
	arrayStack.Put(1002)
	arrayStack.Put(1003)
	fmt.Println("size:", arrayStack.Size())
	fmt.Println("peek:", arrayStack.Peek())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("peek:", arrayStack.Peek())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("size:", arrayStack.Size())
	arrayStack.Put(1005)
	fmt.Println("pop:", arrayStack.Pop())
}
