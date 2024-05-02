package main

import (
	"fmt"
	"shopping/config"
)

func main() {
	config.InitConfig()
	config.InitLogger(config.Conf.User.Http.ServiceName)
	server := NewUserHttpServer()
	server.Run(fmt.Sprintf("%s:%d", config.Conf.User.Http.Host, config.Conf.User.Http.Port))
}
