package main

import (
	"sync"
)

type ArrayQueue struct {
	data []int
	size int
	lock sync.Mutex
}

func (queue *ArrayQueue) In(v int) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	queue.data = append(queue.data, v)
	queue.size++
}

func (queue *ArrayQueue) Out() int {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		panic("empty queue")
	}

	v := queue.data[0]
	newData := make([]int, queue.size-1, queue.size-1)
	for i := 0; i < queue.size-1; i++ {
		newData[i] = queue.data[i+1]
	}
	queue.data = newData
	queue.size--
	return v
}

func (queue *ArrayQueue) Size() int {
	return queue.size
}

func (queue *ArrayQueue) IsEmpty() bool {
	return queue.size == 0
}

// func main() {
// 	queue := new(ArrayQueue)
// 	queue.In(1001)
// 	queue.In(1002)
// 	queue.In(1003)

// 	fmt.Println(queue.Out())
// 	fmt.Println(queue.Out())
// 	queue.In(1004)
// 	fmt.Println(queue.Out())
// }
