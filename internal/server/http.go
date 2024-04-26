package server

import "github.com/gin-gonic/gin"

func NewUserGin() *gin.Engine {
	server := gin.Default()
	server.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "ok",
		})
	})
	return server
}
