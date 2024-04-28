package biz

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	userPb "shopping/api/user"
)

var _ handler = (*UserHandler)(nil)

type UserHandler struct {
	userClient userPb.UserClient
}

func NewUserHandler(userClient userPb.UserClient) *UserHandler {
	return &UserHandler{
		userClient: userClient,
	}
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/user")
	ug.POST("/signUp", u.signUp)
}

func (u *UserHandler) signUp(ctx *gin.Context) {
	type Req struct {
		Name string `json:"name" validate:"required"`
		Tel  string `json:"tel" validate:"required"`
	}
	var req Req
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zap.S().Errorf("[signUp] invalid params: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := u.userClient.CreateUser(ctx, &userPb.CreateUserRequest{
		Name:      req.Name,
		Telephone: req.Tel,
	})
	if err != nil {
		zap.S().Errorf("[signUp] create user failed: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": user,
	})
}
