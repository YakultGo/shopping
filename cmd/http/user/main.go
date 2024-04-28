package main

import (
	"fmt"
	"shopping/config"
)

func main() {
	config.InitConfig()
	config.InitLogger("http-user")
	server := NewUserHttpServer()
	server.Run(fmt.Sprintf("%s:%d", config.Conf.User.Http.Host, config.Conf.User.Http.Port))
}
