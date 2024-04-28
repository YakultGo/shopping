package biz

import "github.com/gin-gonic/gin"

func NewUserGin(userHandler *UserHandler) *gin.Engine {
	server := gin.Default()
	userHandler.RegisterRoutes(server)
	return server
}
