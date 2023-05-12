package benchmark

import (
	"testing"
	"time"
)

func timeConsumingTask() {
	time.Sleep(5000 * time.Millisecond)
}

func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Foo()
	}
}

func BenchmarkBoo1(b *testing.B) {
	timeConsumingTask() // 不相干的操作
	b.ResetTimer()      // 重置计时器
	for i := 0; i < b.N; i++ {
		Foo()
	}
}

func BenchmarkBoo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()       // 停止计时器
		timeConsumingTask() // 不相干的操作
		b.StartTimer()      // 恢复计时器
		Foo()
	}
}
