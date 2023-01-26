package main

import (
	"fmt"
	"net"

	"github.com/lixvyang/betxin-microservice/user/config"
	"github.com/lixvyang/betxin-microservice/user/discovery"
	"github.com/lixvyang/betxin-microservice/user/internal/handler"
	"github.com/lixvyang/betxin-microservice/user/internal/repository"
	service "github.com/lixvyang/betxin-microservice/user/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	config.InitConfig()
	repository.InitDB()

	etcdAddress := []string{viper.GetString("etcd.address")}
	etcdRegister := discovery.NewRegister(etcdAddress, logrus.New())
	grpcAddress := viper.GetString("server.grpcAddress")
	defer etcdRegister.Stop()

	userNode := discovery.Server{
		Name: viper.GetString("server.user"),
		Addr: grpcAddress,
	}
	server := grpc.NewServer()
	defer server.Stop()

	// 绑定service
	service.RegisterUserServiceServer(server, handler.NewUserSerivce())
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}
	if _, err := etcdRegister.Register(userNode, 10); err != nil {
		panic(fmt.Sprintf("start server failed, err: %v", err))
	}
	logrus.Info("server started listen on ", grpcAddress)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}

}
