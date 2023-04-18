package main

// Golang 选项模式（Option 模式）

// Option 模式的专业术语为：Functional Options Pattern（函数式选项模式）
// Option 模式为 golang 的开发者提供了将一个函数的参数设置为可选的功能，
// 也就是说我们可以选择参数中的某几个，并且可以按任意顺序传入参数。
// 比如针对特殊场景需要不同参数的情况，C++ 可以直接用重载来写出任意个同名函数，在任意场景调用的时候使用同一个函数名即可；
// 但同样情况下，在 golang 中我们就必须在不同的场景使用不同的函数，并且参数传递方式可能不同的人写出来是不同的样子，
// 这将导致代码可读性差，维护性差。

// Option模式的优缺点
// 优点：
// 1. 支持传递多个参数，并且在参数个数、类型发生变化时保持兼容性
// 2. 任意顺序传递参数
// 3. 支持默认值
// 4. 方便拓展
// 缺点：
// 1. 增加许多 function，成本增大
// 2. 参数不太复杂时，尽量少用

import (
	"log"
	"time"
)

func main() {
	svr := NewServer(
		WithHost("localhost"),
		WithPort(8080),
		WithTimeout(time.Minute),
		WithMaxConn(120),
	)
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}

type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

func NewServer(options ...func(*Server)) *Server {
	svr := &Server{}
	for _, o := range options {
		o(svr)
	}
	return svr
}

func (s *Server) Start() error {
	// todo
	log.Println("server started.")
	return nil
}

func WithHost(host string) func(*Server) {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) func(*Server) {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxConn(maxConn int) func(*Server) {
	return func(s *Server) {
		s.maxConn = maxConn
	}
}
