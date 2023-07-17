package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-sample/interceptor/pb"
	"log"
	"net"
	"time"
)

func main() {
	// 监听本地端口，作为服务端
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(timer, limiter),
	)

	// 注册 HelloServer
	pb.RegisterHelloServiceServer(server, &HelloServer{})

	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}

type HelloServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloServer) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	time.Sleep(100 * time.Millisecond)

	resp := &pb.HelloResponse{
		Score: int32(len(request.Name)),
	}

	return resp, nil
}

// timer 服务端时间消耗拦截器
func timer(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	log.Println("enter server timer interceptor")
	start := time.Now()
	resp, err := handler(ctx, req)
	log.Printf("time consuming: %d ms", time.Since(start).Milliseconds())
	log.Println("leave server timer interceptor")
	return resp, err
}

var limitChan = make(chan struct{}, 10) // 瞬间并发度限制

// limiter 服务端请求限制拦截器
func limiter(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	log.Println("enter server limiter interceptor")
	limitChan <- struct{}{}
	defer func() {
		<-limitChan
	}()
	resp, err := handler(ctx, req)
	log.Println("leave server limiter interceptor")
	return resp, err
}
