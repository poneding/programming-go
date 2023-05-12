# fuzz testing

模糊测试

## 写法

```go
func FuzzFoo(f *testing.F) int {
    f.Fuzz(testFooFunc)
    return 0
}
```

## 示例

文件 *[foo_v1.go](./foo_v1.go)*

```go
package fuzzing

func ReverseV1(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
```

文件：*[foo_v1_test.go](./foo_v1_test.go)*

```go
package fuzzing

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverseV1(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := ReverseV1(orig)
		doubleRev := ReverseV1(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
```

## 执行

```bash
$ go test -fuzz=FuzzReverseV1 -fuzztime 30s
fuzz: elapsed: 0s, gathering baseline coverage: 0/11 completed
fuzz: elapsed: 0s, gathering baseline coverage: 11/11 completed, now fuzzing with 8 workers
fuzz: minimizing 39-byte failing input file
fuzz: elapsed: 0s, minimizing
--- FAIL: FuzzReverseV1 (0.04s)
    --- FAIL: FuzzReverseV1 (0.00s)
        foo_v1_test.go:20: Reverse produced invalid UTF-8 string "\xbb\xb9\xe3"
    
    Failing input written to testdata/fuzz/FuzzReverseV1/798ff9627bd35788
    To re-run:
    go test -run=FuzzReverseV1/798ff9627bd35788
FAIL
exit status 1
FAIL    fuzzing 0.203s
```

> 测试未通过，因为 `ReverseV1` 不能处理非 UTF-8 字符串。
> 测试结束后，会在目录下生成一个 `testdata` 目录，里面包含了测试失败的输入用例。下次测试会自动加载这些用例，作为回归测试。

## 修复

文件：*[foo_v2.go](./foo_v2.go)*

```go
package fuzzing

import (
	"errors"
	"unicode/utf8"
)

func ReverseV2(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	b := []rune(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}
```

文件：*[foo_v2_test.go](./foo_v2_test.go)*

```go
package fuzzing

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverseV2(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err := ReverseV2(orig)
		if err != nil {
			return
		}
		doubleRev, err := ReverseV2(rev)
		if err != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
```

## 再次执行

```bash
$ rm -r testdata && go test -fuzz=FuzzReverseV2 -fuzztime 30s
fuzz: elapsed: 0s, gathering baseline coverage: 0/51 completed
fuzz: elapsed: 0s, gathering baseline coverage: 51/51 completed, now fuzzing with 8 workers
fuzz: elapsed: 3s, execs: 1063165 (354350/sec), new interesting: 0 (total: 51)
fuzz: elapsed: 6s, execs: 2209719 (382099/sec), new interesting: 0 (total: 51)
fuzz: elapsed: 9s, execs: 3299668 (363426/sec), new interesting: 0 (total: 51)
fuzz: elapsed: 12s, execs: 4373591 (357885/sec), new interesting: 0 (total: 51)
fuzz: elapsed: 15s, execs: 5469249 (365186/sec), new interesting: 0 (total: 51)
fuzz: elapsed: 18s, execs: 6506330 (345800/sec), new interesting: 0 (total: 51)
fuzz: elapsed: 21s, execs: 7680642 (391323/sec), new interesting: 0 (total: 51)
fuzz: elapsed: 24s, execs: 8772135 (363835/sec), new interesting: 0 (total: 51)
fuzz: elapsed: 27s, execs: 9852604 (360204/sec), new interesting: 0 (total: 51)
fuzz: elapsed: 30s, execs: 10942204 (363055/sec), new interesting: 0 (total: 51)
fuzz: elapsed: 30s, execs: 10942204 (0/sec), new interesting: 0 (total: 51)
PASS
ok      fuzzing 30.670s
```

## 参考

- [Go Fuzzing](https://go.dev/security/fuzz/)
- [Tutorial: Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz)
