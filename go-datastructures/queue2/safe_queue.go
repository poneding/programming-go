package queue2

import "sync"

type SafeQueue struct {
	sync.Mutex
	q *Queue
}

func NewSafeQueue() *SafeQueue {
	return &SafeQueue{
		q: NewQueue(),
	}
}

func (s *SafeQueue) Len() int {
	s.Lock()
	defer s.Unlock()
	return s.q.Len()
}

func (s *SafeQueue) Push(v interface{}) {
	s.Lock()
	defer s.Unlock()
	s.q.Push(v)
}

func (s *SafeQueue) Pop() (interface{}, bool) {
	s.Lock()
	defer s.Unlock()
	return s.q.Pop()
}

func (s *SafeQueue) Front() (interface{}, bool) {
	s.Lock()
	defer s.Unlock()
	return s.q.Front()
}
