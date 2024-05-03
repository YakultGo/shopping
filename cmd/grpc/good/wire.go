//go:build wireinject

package main

import (
	"github.com/google/wire"
	"shopping/internal/server"
)

func NewGoodGrpc() *server.GoodServer {
	wire.Build(
		server.NewGoodServer,
	)
	return new(server.GoodServer)
}
