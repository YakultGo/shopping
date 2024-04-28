package register

import (
	"fmt"
	"os"
	"os/signal"
	"shopping/config"
	"syscall"

	"github.com/hashicorp/consul/api"
)

type RegistryClient interface {
	Register(address string, port int, name string, tags []string, id string) error
	DeRegister(serviceId string) error
	GracefulStop(serviceId string) error
}
type Registry struct {
	Host string
	Port int
}

func NewRegistryClient(host string, port int) RegistryClient {
	return &Registry{
		Host: host,
		Port: port,
	}
}

// Register 服务注册
func (r *Registry) Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", config.Conf.Consul.Host, config.Conf.Consul.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}
	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	// 生成对应grpc的检查对象
	check := &api.AgentServiceCheck{
		// 这里的address要写host.docker.internal，不能直接写localhost
		GRPC:                           fmt.Sprintf("%s:%d", "host.docker.internal", port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	registration.Check = check
	err = client.Agent().ServiceRegister(registration)
	return err
}

func (r *Registry) DeRegister(serviceId string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", config.Conf.Consul.Host, config.Conf.Consul.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}
	err = client.Agent().ServiceDeregister(serviceId)
	return err
}

func (r *Registry) GracefulStop(serviceId string) error {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-quit
	err := r.DeRegister(serviceId)
	return err
}
