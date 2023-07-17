package main

// 1. 不使用 wire 时，需要手动初始化依赖
//func main() {
//	message := NewMessage()
//	greeter := NewGreeter(message)
//	event := NewEvent(greeter)
//
//	event.Start()
//}

// 2. 使用 wire：
// 2.1. 在 wire.go 中 使用 wire.Build() 函数来初始化依赖
// 2.2. 在 main.go 中 调用 InitializeEvent() 函数来初始化依赖
// 2.3. 使用命令 wire gen ./... 来生成 wire_gen.go 文件
// 2.4. 执行 go run main.go 来运行程序
func main() {
	event := InitializeEvent()
	event.Start()
}
