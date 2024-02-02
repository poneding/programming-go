package concurrentmap

import (
	"errors"
	"sync"
	"time"
)

// 实现一个 map
// 要求：1. 面向高并发
//  2. 只存在插入和查询操作 O(1)
//  3. 查询时，存在直接返回，不存在则阻塞，直到有插入操作插入数据，此时返回插入的数据，如果超出一定时间，则返回错误
type concurrentMap struct {
	// 互斥锁
	sync.Mutex
	// 存储数据的 map
	m map[any]any
	// 当有插入操作时，通知查询操作
	keyc map[any]chan struct{}
}

func NewConcurrentMap() *concurrentMap {
	return &concurrentMap{
		m: make(map[any]any),
	}
}

func (m *concurrentMap) Put(key, value any) {
	m.Lock()
	defer m.Unlock()
	m.m[key] = value

	if ch, ok := m.keyc[key]; ok {
		// 说明有查询操作在等待，关闭通道通知查询操作
		close(ch) // 关闭 chan 会通知所有等待的 goroutine，会读到零值
		delete(m.keyc, key)
	}
}

func (m *concurrentMap) Get(key any, timeout time.Duration) (value any, err error) {
	m.Lock()
	if v, ok := m.m[key]; ok {
		m.Unlock()
		return v, nil
	}

	// 说明不存在，需要等待
	if m.keyc[key] == nil {
		m.keyc[key] = make(chan struct{})
	}

	m.Unlock()

	select {
	case <-m.keyc[key]:
		// 说明有插入操作，那么则表明 m.m[key] 已经有值了
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}

	m.Lock() // 再次加锁，读取数据
	v := m.m[key]
	m.Unlock()
	return v, nil
}
