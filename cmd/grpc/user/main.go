package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	userPb "shopping/api/user"
	"shopping/config"
	"shopping/internal/server"
	"shopping/pkg/register"
	"shopping/pkg/util"
)

func main() {
	config.InitConfig()
	config.InitLogger("grpc-user")
	port, err := util.GetFreePort()
	if err != nil {
		zap.S().Error(err)
	}
	grpcServer := grpc.NewServer()
	userPb.RegisterUserServer(grpcServer, server.NewUserServer())
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Conf.User.Grpc.Host, port))
	if err != nil {
		zap.S().Error(err)
	}
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			zap.S().Error(err)
		}
	}()
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	// 注册user服务
	registerClient := register.NewRegistryClient(config.Conf.Consul.Host, config.Conf.Consul.Port)
	ServiceId := fmt.Sprintf("%s", uuid.NewV4())
	err = registerClient.Register(
		config.Conf.User.Grpc.Host,
		port,
		config.Conf.User.Grpc.ServiceName,
		config.Conf.User.Grpc.Tags,
		ServiceId,
	)
	if err != nil {
		zap.S().Errorf("register user to consul failed, err:%v", err)
	}
	// grpc服务启动成功
	zap.S().Infof("grpc user server start success, port: %d", port)
	// 优雅关闭
	err = registerClient.GracefulStop(ServiceId)
	if err != nil {
		zap.S().Error("user服务注销失败: ", err.Error())
	}
	zap.S().Infof("user服务注销成功")
}
