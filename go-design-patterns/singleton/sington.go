package main

import (
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type singleton struct {
	ID string
}

var (
	instance    *singleton
	initialized uint32
	mu          sync.Mutex

	instance2 *singleton
	once      sync.Once
)

func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{
			ID: strconv.FormatInt(time.Now().UnixNano(), 10),
		}
	}
	return instance
}

func Instance2() *singleton {
	once.Do(func() {
		instance2 = &singleton{
			ID: strconv.FormatInt(time.Now().UnixNano(), 10),
		}
	})
	return instance2
}
