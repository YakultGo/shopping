package connect

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	userPb "shopping/api/user"
	"shopping/config"
)

func NewUserGrpc() userPb.UserClient {
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s",
			config.Conf.Consul.Host,
			config.Conf.Consul.Port,
			config.Conf.User.Grpc.ServiceName),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Errorf("[NewUserGrpc] grpc dial err: %v", err)
		return nil
	}
	return userPb.NewUserClient(userConn)
}
