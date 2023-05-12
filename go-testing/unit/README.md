# unit testing

单元测试

## 写法

```go
func TestFoo(t *testing.T) {
    Foo()
}
```

函数名以 `Test` 前缀开头。`*testing.T` 类型的参数 `t` 用于报告测试失败和附加日志信息。`Foo()` 是被测试的函数。

## 示例

文件 *foo.go*

```go
package unit

func Reverse(s string) string {
	b := []rune(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
```

文件：*foo_test.go*

```go
package unit

import (
	"testing"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev := Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}
```

## 执行

```bash
$ go test -v
=== RUN   TestReverse
--- PASS: TestReverse (0.00s)
PASS
ok      unit    0.349s
```
