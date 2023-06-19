package main

import (
	"context"
	"errors"
	"fmt"
	"grpc-go/pb"
	"log"
	"net"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	errMissingMetadata = status.Error(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Error(codes.Unauthenticated, "invalid token")
)

type server struct {
	// 嵌套了一个未实现的接口，这样就可以不用实现接口中的所有方法
	pb.UnimplementedCreateUserServiceServer
}

// 实现 CreateUserServiceServer 的 CreateUser 方法
func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if req.FirstName == "" || req.LastName == "" {
		return nil, fmt.Errorf("first name or last name is empty: %w", errors.New("invalid request"))
	}
	return &pb.CreateUserResponse{
		ID:        uuid.New().ID(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
	}, nil
}

func main() {
	// 1. 创建一个 TCP 监听器
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	// 2. 创建一个 gRPC 服务器
	// s := grpc.NewServer()
	s := grpc.NewServer(setupServerOptions()...)

	// 3. 在 gRPC 服务器上注册 CreateUserServiceServer
	pb.RegisterCreateUserServiceServer(s, &server{})

	log.Println("gRPC server is running on port 8080")
	// 4. 在监听器上启动 gRPC 服务器
	if s.Serve(l); err != nil {
		log.Fatalln(err)
	}
}

// 拦截器
func setupServerOptions() []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.UnaryInterceptor(unaryInterceptor),
	}
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	metadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	authorization := metadata.Get("authorization")
	if len(authorization) == 0 {
		return nil, errInvalidToken
	}
	if token := strings.Trim(authorization[0], "Bearer "); token != "123456" {
		return nil, errInvalidToken
	}

	i, err := handler(ctx, req)
	// if err != nil {
	// 	return nil, err
	// }
	// return i, nil
	return i, err
}
