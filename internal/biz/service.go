package biz

import (
	"github.com/gin-gonic/gin"
)

func NewUserGin(userHandler *UserHandler, handlerFunc []gin.HandlerFunc) *gin.Engine {
	server := gin.Default()
	server.Use(handlerFunc...)
	userHandler.RegisterRoutes(server)
	return server
}
