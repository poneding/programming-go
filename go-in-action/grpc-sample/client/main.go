package main

import (
	"context"
	"grpc-sample/pb"
	"io"
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
	// createUserClient := pb.NewCreateUserServiceClient(conn)
	// callCreateUser(createUserClient)

	// 3. 创建一个 GreetServiceClient 客户端
	greetClient := pb.NewGreetServiceClient(conn)
	// callSayHello(greetClient)
	// callSayHelloServerStreaming(greetClient)
	// callSayHelloClientStreaming(greetClient)
	callSayHelloBidirectionalStreaming(greetClient)
}

func callCreateUser(client pb.CreateUserServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", "Bearer 123456"))
	resp, err := client.CreateUser(ctx, &pb.CreateUserRequest{
		FirstName: "Pone",
		LastName:  "Ding",
		Age:       30,
	})
	if err != nil {
		panic(err)
	}

	log.Println("create user successfully, User ip:", resp.ID)
}

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", "Bearer 123456"))

	resp, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		panic(err)
	}

	log.Println("say hello successfully, Message:", resp.Message)
}

func callSayHelloServerStreaming(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", "Bearer 123456"))

	stream, err := client.SayHelloServerStreaming(ctx, &pb.NamesList{
		Names: []string{"Pone", "Jay", "Tom"},
	})
	if err != nil {
		panic(err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		log.Println("say hello successfully, Message:", resp.Message)
	}
}

func callSayHelloClientStreaming(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", "Bearer 123456"))

	namesList := pb.NamesList{
		Names: []string{"Pone", "Jay", "Tom"},
	}

	stream, err := client.SayHelloClientStreaming(ctx)
	if err != nil {
		panic(err)
	}

	for _, name := range namesList.Names {
		if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
			panic(err)
		}
		log.Println("send name:", name)
		time.Sleep(time.Second)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}

	log.Println("say hello successfully, Message:", resp.Messages)
}

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", "Bearer 123456"))

	namesList := pb.NamesList{
		Names: []string{"Pone", "Jay", "Tom"},
	}

	stream, err := client.SayHelloBidirectionalStreaming(ctx)
	if err != nil {
		panic(err)
	}

	waitC := make(chan struct{})
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			log.Println("receive message:", message.Message)
		}
		close(waitC)
	}()

	for _, name := range namesList.Names {
		if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
			panic(err)
		}
		log.Println("send name:", name)
		time.Sleep(time.Second)
	}

	stream.CloseSend()
	<-waitC
	log.Println("SayHelloBidirectionalStreaming completed")
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
