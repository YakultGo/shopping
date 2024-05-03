//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"shopping/config"
	"shopping/internal/biz"
	"shopping/internal/middlewares"
	"shopping/internal/middlewares/jwt"
	"shopping/pkg/consul/connect"
)

func NewUserHttpServer() *gin.Engine {
	wire.Build(
		connect.NewUserGrpc,
		connect.NewSmsGrpc,
		biz.NewUserHandler,
		biz.NewGoodHandler,
		config.NewRedis,
		jwt.NewRedisJWTHandler,
		middlewares.NewMiddlewares,
		biz.NewUserGin,
	)
	return new(gin.Engine)
}
