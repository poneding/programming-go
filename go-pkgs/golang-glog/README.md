# github.com/golang/glog

[`golang/glog`](https://pkg.go.dev/github.com/golang/glog) 可以按照等级将日志打印到标准错误输出或文件。

## 引入

```go
import "github.com/golang/glog"
```

## 安装

```bash
go get github.com/golang/glog
```

## 示例

```go
package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	// 
	flag.Parse()
	defer glog.Flush()

	glog.Infoln("info") // 
	glog.V(1).Infoln("v1 info")
	glog.V(2).Infoln("v2 info")
	glog.V(2).Infoln("v3 info")

	if glog.V(2) {
		// 设置了 -v=2，将会打印
		glog.Infoln("v2 info 1")
	}
	glog.Warningln("warning")
	glog.Errorln("error")
	// glog.Fatalln("fatal") // 将调用 os.Exit(2)
	// glog.Exitln("exit") // 将调用 os.Exit(1)

	// fmt.Println("will not print")
}
```

## 运行

打印到标准错误输出：

```bash
go run main.go -logtostderr -v=2
```

打印到文件：

```bash
go run main.go -log_dir=./logs -v=2
```

同时打印到标准错误输出和文件：

```bash
go run main.go -log_dir=./logs -alsologtostderr -v=2
```

> ⚠️ 注意：  
> 1、打印到文件需要确保 `log_dir` 目录存在；  
> 2、如果使用 `logtostderr`，则 `log_dir` 无效（不会打印到文件）。  
