package main

import (
	"sync"
	"time"
)

// 实现阻塞读且并发安全的map
// func main() {

// }

type sp interface {
	Put(key string, val interface{})
	Get(key string, timeout time.Duration) interface{}
}

type MyMap struct {
	c   map[string]*entry
	rmx *sync.RWMutex
}

type entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
}

func (m *MyMap) Put(key string, val interface{}) {
	m.rmx.Lock()
	defer m.rmx.Unlock()

	item, ok := m.c[key]

	if !ok {
		m.c[key] = &entry{
			value:   val,
			isExist: true,
		}
		return
	}
	item.value = val
	if !item.isExist {
		if item.ch != nil {
			close(item.ch)
			item.ch = nil
		}
	}
	return
}
