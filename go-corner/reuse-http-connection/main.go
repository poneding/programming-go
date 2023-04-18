package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"
)

func main() {
	clientTrace := httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) {
			fmt.Println("conn is reused:", info.Reused)
		},
	}
	traceCtx := httptrace.WithClientTrace(context.Background(), &clientTrace)

	req, err := http.NewRequestWithContext(traceCtx, http.MethodGet, "http://example.com", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	// !!! 如果想要复用连接，必须先将 resp.Body 读取完全，然后在关闭 resp.Body
	// 否则，下次请求时，会新建连接
	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		panic(err)
	}
	resp.Body.Close()

	req2, err := http.NewRequestWithContext(traceCtx, http.MethodGet, "http://example.com", nil)
	if err != nil {
		panic(err)
	}

	resp2, err := http.DefaultClient.Do(req2)
	if err != nil {
		panic(err)
	}
	resp2.Body.Close()
}

/*
$ go run reuse-http-connection.go
conn is reused: false
conn is reused: true
*/
