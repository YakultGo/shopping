//go:build wireinject

package main

import (
	"github.com/google/wire"
	"shopping/internal/server"
)

func NewUserGrpc() *server.UserServer {
	wire.Build(
		server.NewUserServer,
	)
	return new(server.UserServer)
}
