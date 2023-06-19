package main

import (
	"context"
	"grpc-go/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	// 1. 创建一个 gRPC 连接
	conn, err := grpc.Dial("localhost:8080", setupDialOptions()...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 2. 创建一个 CreateUserServiceClient 客户端
	client := pb.NewCreateUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", "Bearer 123456"))
	resp, err := client.CreateUser(ctx, &pb.CreateUserRequest{
		FirstName: "Peng",
		LastName:  "Ding",
		Age:       30,
	})
	if err != nil {
		panic(err)
	}

	log.Println("create user successfully, User ID:", resp.ID)
}

func setupDialOptions() []grpc.DialOption {
	return []grpc.DialOption{
		// tls 加密参考示例：https://github.com/grpc/grpc-go/blob/master/examples/features/authentication/server/main.go
		// grpc.WithPerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
		// 	AccessToken: "123456",
		// })), // 使用 oauth 认证, 必须需要证书
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
}
