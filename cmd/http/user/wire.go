//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"shopping/internal/biz"
	"shopping/pkg/consul/connect"
)

func NewUserHttpServer() *gin.Engine {
	wire.Build(
		biz.NewUserHandler,
		connect.NewUserGrpc,
		biz.NewUserGin,
	)
	return new(gin.Engine)
}
