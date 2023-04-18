package main

import (
	"sync"
)

type LinkNode struct {
	v    int
	next *LinkNode
}

type LinkQueue struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

func (queue *LinkQueue) In(v int) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		queue.root = &LinkNode{v: v}
	} else {
		n := queue.root
		for n.next != nil {
			n = n.next
		}
		n.next = &LinkNode{v: v}
	}
	queue.size++
}

func (queue *LinkQueue) Out() int {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		panic("empty queue")
	}

	v := queue.root.v
	queue.root = queue.root.next
	queue.size--
	return v
}

func (queue *LinkQueue) Size() int {
	return queue.size
}

func (queue *LinkQueue) IsEmpty() bool {
	return queue.size == 0
}

// func main() {
// 	queue := new(LinkQueue)
// 	queue.In(1001)
// 	queue.In(1002)
// 	queue.In(1003)

// 	fmt.Println(queue.Out())
// 	fmt.Println(queue.Out())
// 	queue.In(1004)
// 	fmt.Println(queue.Out())
// }
