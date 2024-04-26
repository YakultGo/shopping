//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"shopping/internal/server"
)

func NewUserHttpServer() *gin.Engine {
	wire.Build(
		server.NewUserGin,
	)
	return new(gin.Engine)
}
