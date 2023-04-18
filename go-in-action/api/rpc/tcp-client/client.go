package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1000")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reps string
	err = client.Call("HelloService.Hello", "Jay Chou", &reps)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reps)
}
