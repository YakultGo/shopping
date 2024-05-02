//go:build wireinject

package main

import (
	"github.com/google/wire"
	"shopping/config"
	"shopping/internal/server"
)

func NewSmsGrpc() *server.SmsServer {
	wire.Build(
		config.NewRedis,
		server.NewSmsServer,
	)
	return new(server.SmsServer)
}
