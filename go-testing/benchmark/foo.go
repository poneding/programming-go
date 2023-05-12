package benchmark

import (
	"bytes"
	"time"

	"github.com/google/uuid"
)

func Foo() string {
	// 模拟加载内存
	data := make([]byte, 5*1024*1024)
	var buf bytes.Buffer
	buf.Write(data)
	time.Sleep(100 * time.Millisecond)
	return uuid.NewString()
}
