package main

import (
	"errors"
	"fmt"
)

func main() {
	q := new(Queue)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println("size:", q.size)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println("size:", q.size)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println("size:", q.size)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println("size:", q.size)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
}

// 实现一个队列 1.5级
// 使用数组实现一个队列Queue，队列是一个先进先出的数据结构。
// 这个队列需要具有以下几个方法
// Push(v int) -- 将一个元素放入队列的尾部。
// Pop() bool -- 从队列首部移除元素。
// Peek() (int, error)  -- 返回队列首部的元素。
// Empty() bool -- 返回队列是否为空。
type Queue struct {
	data []int
	size int
}

func (q *Queue) Push(v int) {
	q.data = append(q.data, v)
	q.size++
}

func (q *Queue) Pop() bool {
	if q.size == 0 {
		return false
	}
	for i := 0; i < q.size-1; i++ {
		q.data[i] = q.data[i+1]
	}
	q.size--
	return true
}

func (q *Queue) Peek() (int, error) {
	if q.size == 0 {
		return 0, errors.New("no data")
	}
	return q.data[0], nil
}

func (q *Queue) Empty() bool {
	return q.size == 0
}
