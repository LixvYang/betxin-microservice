package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lixvyang/betxin-microservice/api-gateway/config"
	"github.com/lixvyang/betxin-microservice/api-gateway/discovery"
	"github.com/lixvyang/betxin-microservice/api-gateway/internal/service"
	"github.com/lixvyang/betxin-microservice/api-gateway/router"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

func main() {
	config.InitConfig()
	etcdAddress := []string{viper.GetString("etcd.address")}
	etcdRegister := discovery.NewResolver(etcdAddress, logrus.New())
	resolver.Register(etcdRegister)
	go startListen()
	{
		osSignal := make(chan os.Signal, 1)
		signal.Notify(osSignal, os.Interrupt, syscall.SIGTERM, syscall.SIGTERM, syscall.SIGINT, syscall.SIGTERM)
		s := <-osSignal
		fmt.Println("exit...", s)
	}
	fmt.Println("gateway listen on :3000")
}

func startListen() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	userConn, err := grpc.Dial(viper.GetString("server.grpcAddress"), opts...)
	if err != nil {
		panic("userConn error!")
	}
	userService := service.NewUserServiceClient(userConn)

	ginRouter := router.NewRouter(userService)
	server := &http.Server{
		Addr:           viper.GetString("server.HttpPort"),
		Handler:        ginRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("绑定失败：端口被占用 %d", err)
	}
}
