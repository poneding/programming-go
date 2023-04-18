package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

// Benchmark
func BenchmarkMutexForMultipleReaders(b *testing.B) {
	var (
		lastValue uint64
		mux       sync.RWMutex
		wg        sync.WaitGroup
	)

	conf := &Config{
		a: []int{0, 0, 0, 0, 0},
	}

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < b.N; i++ {
				mux.RLock()
				atomic.SwapUint64(&lastValue, uint64(conf.a[0]))
				mux.RUnlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
