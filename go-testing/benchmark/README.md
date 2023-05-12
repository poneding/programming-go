# benchmark testing

基准测试

## 写法

```go
func BenchmarkFoo(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Foo()
    }
}
```

函数名以 `Benchmark` 前缀开头。在循环 `for` 中调用被测函数 `Foo()` 。`b.N` 表示可变的迭代次数。运行基准测试时，Go 会尝试使其与请求的基准测试时间相匹配。基准时间默认设置为 1 秒，可以通过 `-benchtime` 标志更改。`b.N` 从 1 开始；如果基准测试在 1 秒内完成，则 `b.N` 会增加，并且基准测试会再次运行，直到 `b.N` 大致匹配 `benchtime`。

## 示例

文件 *[foo.go](./foo.go)*

```go
package benchmark

import (
	"bytes"

	"github.com/google/uuid"
)

func Foo() string {
	// 模拟加载内存
	data := make([]byte, 5*1024*1024)
	var buf bytes.Buffer
	buf.Write(data)
	return uuid.NewString()
}
```

文件：*[foo_test.go](./foo_test.go)*

```go
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

func BenchmarkFoo1(b *testing.B) {
	timeConsumingTask() // 不相干的操作
	b.ResetTimer()      // 重置计时器
	for i := 0; i < b.N; i++ {
		Foo()
	}
}

func BenchmarkFoo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()       // 停止计时器
		timeConsumingTask() // 不相干的操作
		b.StartTimer()      // 恢复计时器
		Foo()
	}
}
```

## 执行

```bash
$ go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: benchmark
BenchmarkFoo-8                10         103129833 ns/op        10485862 B/op          4 allocs/op
BenchmarkFoo1-8               10         103937912 ns/op        10485837 B/op          4 allocs/op
BenchmarkFoo2-8               10         102841421 ns/op        10485844 B/op          4 allocs/op
PASS
ok      benchmark       68.869s
```

> `-benchmem` 标志会在报告中包含内存分配的统计数据。
> 这里，`BenchmarkFoo` 基准测试时间为 1s，`b.N` 为 10，平均每次迭代耗时 103129833 ns，每次迭代分配 10485862 B 内存，每次迭代有 4 次内存分配。
> 可以通过 `-benchtime` 标志更改基准测试时间，例如：`go test -bench=BenchmarkFoo -benchtime=10s`；
> 可以通过 `-count` 标志指定每个基准测试运行的次数，例如：`go test -bench=BenchmarkFoo -count=10`；
