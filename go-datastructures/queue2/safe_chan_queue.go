package queue2

import (
	"sync"
	"time"
)

const (
	signalInterval = 200
	signalChanSize = 10
)

type SafeChanQueue struct {
	sync.Mutex
	q *Queue
	C chan struct{}
}

func NewSafeChanQueue() *SafeChanQueue {
	scq := &SafeChanQueue{
		q: NewQueue(),
		C: make(chan struct{}, signalChanSize),
	}

	go func() {
		ticker := time.NewTicker(time.Millisecond * signalInterval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if scq.q.Len() > 0 {
					select {
					case scq.C <- struct{}{}:
						// signaled
					default:
						// not block this goroutine
					}
				}
			}
		}
	}()
	return scq
}

func (s *SafeChanQueue) Len() int {
	s.Lock()
	defer s.Unlock()
	return s.q.Len()
}

func (s *SafeChanQueue) Push(v interface{}) {
	s.Lock()
	defer s.Unlock()
	s.q.Push(v)
}

func (s *SafeChanQueue) Pop() (interface{}, bool) {
	s.Lock()
	defer s.Unlock()
	return s.q.Pop()
}

func (s *SafeChanQueue) Front() (interface{}, bool) {
	s.Lock()
	defer s.Unlock()
	return s.q.Front()
}
