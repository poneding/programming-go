package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()       // 使用 flag.Parse() 方法解析 glog 中的命令行参数。
	defer glog.Flush() // 使用 glog.Flush() 将缓冲区的内容输出。

	glog.Infoln("info")

	// 只会打印 <= {v} 的日志
	// 如果运行时设置 -v=2，将会打印 <= 2 的日志
	// 因此 glog.V(3) 的日志不会打印
	glog.V(1).Infoln("v1 info")
	glog.V(2).Infoln("v2 info")
	glog.V(3).Infoln("v3 info")

	if glog.V(2) {
		// 设置了 -v=2，将会打印
		glog.Infoln("v2 info embedded")
	}
	glog.Warningln("warning")
	glog.Errorln("error")
	// glog.Fatalln("fatal") // 将调用 os.Exit(2)
	// glog.Exitln("exit")   // 将调用 os.Exit(1)
}
