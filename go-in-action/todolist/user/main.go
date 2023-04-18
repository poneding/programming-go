package main

import (
	"fmt"
	"net"
	"user/config"
	"user/discovery"
	"user/internal/handler"
	"user/internal/repository"

	service "user/internal/service/pb"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	config.Init()
	repository.InitDB()

	// etcd
	etcdEndpoints := viper.GetStringSlice("etcd.endpoints")

	etcdRegister := discovery.NewRegister(etcdEndpoints, logrus.New())
	defer etcdRegister.Stop()
	grpcAddress := viper.GetString("server.grpcAddress")

	server := grpc.NewServer()
	defer server.Stop()

	service.RegisterUserServiceServer(server, handler.NewUserService())

	listen, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}

	if _, err := etcdRegister.Register(discovery.Server{
		Name:    viper.GetString("server.domain"),
		Address: grpcAddress,
	}, 10); err != nil {
		panic(fmt.Sprintf("start server failed, err: %s", err.Error()))
	}

	if err := server.Serve(listen); err != nil {
		logrus.Fatalln(err)
	}
}
