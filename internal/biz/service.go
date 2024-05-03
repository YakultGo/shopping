package biz

import (
	"github.com/gin-gonic/gin"
)

func NewUserGin(userHandler *UserHandler, goodHandler *GoodHandler, handlerFunc []gin.HandlerFunc) *gin.Engine {
	server := gin.Default()
	server.Use(handlerFunc...)
	userHandler.RegisterRoutes(server)
	goodHandler.RegisterRoutes(server)
	return server
}
