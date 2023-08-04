package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	// 模拟另外一个线程，创建目录和文件
	go func() {
		var i int
		for i <= 100000 {
			os.Create("testdata/dir1/" + fmt.Sprintf("cfile%d", i))
		}
	}()
	// 使用 os.RemoveAll 删除目录时，如果目录下文件被其他线程使用，可能会报错：directory not empty
	if err := os.RemoveAll("testdata/"); err != nil {
		if errors.Is(err, syscall.ENOTEMPTY) {
			fmt.Println("err:", err)
		} else {
			fmt.Println(err)
		}
	}
	time.Sleep(2)
}
