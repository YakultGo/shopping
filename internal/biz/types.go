package biz

import "github.com/gin-gonic/gin"

type handler interface {
	RegisterRoutes(server *gin.Engine)
}
