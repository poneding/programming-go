package main

import (
	"context"
	"grpc-go/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1. 创建一个 gRPC 连接
	// conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure()) // grpc.WithInsecure() 已经被弃用 Deprecated
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 2. 创建一个 CreateUserServiceClient 客户端
	client := pb.NewCreateUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
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
