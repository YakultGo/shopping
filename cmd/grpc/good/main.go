package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	goodPb "shopping/api/good"
	"shopping/config"
	"shopping/internal/data/query"
	"shopping/pkg/consul/register"
	"shopping/pkg/util"
)

func main() {
	config.InitConfig()
	config.InitLogger(config.Conf.Good.Grpc.ServiceName)
	// 初始化mysql
	query.SetDefault(config.NewMysql())

	port, err := util.GetFreePort()
	if err != nil {
		zap.S().Error(err)
	}
	grpcServer := grpc.NewServer()
	goodPb.RegisterGoodServer(grpcServer, NewGoodGrpc())
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Conf.Good.Grpc.Host, port))
	if err != nil {
		zap.S().Error(err)
	}
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			zap.S().Error(err)
		}
	}()
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	// 注册Good服务
	registerClient := register.NewRegistryClient(config.Conf.Consul.Host, config.Conf.Consul.Port)
	ServiceId := fmt.Sprintf("%s", uuid.NewV4())
	err = registerClient.Register(
		config.Conf.Good.Grpc.Host,
		port,
		config.Conf.Good.Grpc.ServiceName,
		config.Conf.Good.Grpc.Tags,
		ServiceId,
	)
	if err != nil {
		zap.S().Errorf("register Good to consul failed, err:%v", err)
	}
	// grpc服务启动成功
	zap.S().Infof("grpc Good server start success, port: %d", port)
	// 优雅关闭
	err = registerClient.GracefulStop(ServiceId)
	if err != nil {
		zap.S().Error("Good服务注销失败: ", err.Error())
	}
	zap.S().Infof("Good服务注销成功")
}
