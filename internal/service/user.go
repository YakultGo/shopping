package service

import (
	"github.com/gin-gonic/gin"
)

var _ handler = (*UserHandler)(nil)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/user")
	ug.GET("/signUp", u.signUp)
}

func (u *UserHandler) signUp(ctx *gin.Context) {
	type Request struct {
	}
}
