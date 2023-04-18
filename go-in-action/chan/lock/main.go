package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// m1()
	// m2()
	m3()
}

type Config struct {
	a []int
}

func m1() {
	conf := &Config{} // 读取的时候读出来的是地址，数据有可能被其他goroutine改变，
	// conf := Config{} // 读取时候读出来的是一整块数据，数据不会改变

	go func() {
		i := 0
		for {
			conf.a = []int{i, i + 1, i + 2, i + 3, i + 4}
			i++
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				fmt.Println(conf)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

// 使用sync.Mutex解决读写冲突，但是性能一下子就下去了
func m2() {
	conf := &Config{} // 读取的时候读出来的是地址，数据有可能被其他goroutine改变，
	// conf := Config{} // 读取时候读出来的是一整块数据，数据不会改变

	var lock sync.RWMutex

	go func() {
		i := 0
		for {
			lock.Lock()
			conf.a = []int{i, i + 1, i + 2, i + 3, i + 4}
			lock.Unlock()
			i++
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				lock.RLock()
				fmt.Println(conf)
				lock.RUnlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

// 使用解决读写冲突
func m3() {
	var v atomic.Value

	go func() {
		i := 0
		for {
			conf := &Config{
				a: []int{i, i + 1, i + 2, i + 3, i + 4},
			}
			v.Store(conf)
			i++
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				conf := v.Load()
				fmt.Println(conf)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
