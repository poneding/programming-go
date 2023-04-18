package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (p *HelloService) Hello(req string, resp *string) error {
	*resp = fmt.Sprintf("Hello %s!", req)
	return nil
}

var flagPort = flag.Int("port", 1000, "listen port")

func main() {
	flag.Parse()

	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *flagPort))
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

	// console test:
	// echo -e '{"method":"HelloService.Hello","params":["Jay Chou"],"id":1}' | nc localhost 1000
}
