//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"shopping/internal/service"
)

func NewUserHttpServer() *gin.Engine {
	wire.Build(
		service.NewUserGin,
	)
	return new(gin.Engine)
}
