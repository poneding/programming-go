package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-sample/interceptor/pb"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(":8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// client interceptors
		grpc.WithChainUnaryInterceptor(timer, limiter),
	)
	if err != nil {
		panic(err)
	}

	client := pb.NewHelloServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Pone Ding"})
	if err != nil {
		panic(err)
	}

	log.Printf("resp.Score: %d", resp.Score)
}

// timer 客户端时间消耗拦截器
func timer(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Println("enter client timer interceptor")
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("time consuming: %d ms", time.Since(start).Milliseconds())
	log.Println("leave client timer interceptor")
	return err
}

var limitChan = make(chan struct{}, 10) // 瞬间并发度限制

// limiter 客户端请求限制拦截器
func limiter(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Println("enter client limiter interceptor")
	limitChan <- struct{}{}
	err := invoker(ctx, method, req, reply, cc, opts...)
	<-limitChan
	log.Println("leave client limiter interceptor")
	return err
}
