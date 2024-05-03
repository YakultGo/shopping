package biz

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	smsPb "shopping/api/sms"
	userPb "shopping/api/user"
	"shopping/internal/middlewares/jwt"
	"shopping/pkg/encode"
)

var _ handler = (*UserHandler)(nil)

type UserHandler struct {
	userClient userPb.UserClient
	smsClient  smsPb.SmsClient
	jwtHandler jwt.Handler
}

func NewUserHandler(userClient userPb.UserClient, smsClient smsPb.SmsClient, hdl jwt.Handler) *UserHandler {
	return &UserHandler{
		userClient: userClient,
		smsClient:  smsClient,
		jwtHandler: hdl,
	}
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/user")
	ug.POST("/signUp", u.signUp)
	ug.POST("/sendCode", u.sendCode)
	ug.POST("/login", u.login)
	ug.POST("/logout", u.logout)
	ug.POST("/modify", u.addAddress)
}

func (u *UserHandler) signUp(ctx *gin.Context) {
	type Req struct {
		Name       string `json:"name" validate:"required"`
		Tel        string `json:"tel" validate:"required"`
		Password   string `json:"password" validate:"required"`
		RePassword string `json:"re_password" validate:"required, eqfield=Password"`
	}
	var req Req
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zap.S().Errorf("[signUp] invalid params: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 对密码进行加密
	encodePwd, err := encode.Encode(req.Password)
	if err != nil {
		zap.S().Errorf("[signUp] encode password failed: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user, err := u.userClient.CreateUser(ctx, &userPb.CreateUserRequest{
		Name:      req.Name,
		Telephone: req.Tel,
		Password:  encodePwd,
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

func (u *UserHandler) sendCode(ctx *gin.Context) {
	type Req struct {
		Tel string `json:"tel" validate:"required, mobile"`
	}
	var req Req
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zap.S().Errorf("[sendCode] invalid params: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	code, err := u.smsClient.SendCode(ctx, &smsPb.SendCodeRequest{
		Phone: req.Tel,
		Biz:   "login",
	})
	if err != nil {
		zap.S().Errorf("[sendCode] send code failed: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"短信验证码": code,
	})
}

func (u *UserHandler) login(ctx *gin.Context) {
	type Req struct {
		Name     string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	var req Req
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zap.S().Errorf("[login] invalid params: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 先查用户密码
	user, err := u.userClient.GetUser(ctx, &userPb.GetUserRequest{
		Name: req.Name,
	})
	if err != nil {
		zap.S().Errorf("[login] get user failed: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 对比密码
	if ok := encode.ComparePasswords(user.Password, req.Password); !ok {
		zap.S().Errorf("[login] compare password failed")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "密码错误"})
		return
	}
	// 在这里使用JWT生成token
	if err = u.jwtHandler.SetLoginToken(ctx, user.Id); err != nil {
		zap.S().Errorf("[login] set login token failed: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}

func (u *UserHandler) logout(ctx *gin.Context) {
	err := u.jwtHandler.ClearToken(ctx)
	if err != nil {
		zap.S().Errorf("[logout] clear token failed: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "退出成功",
	})
}

func (u *UserHandler) addAddress(ctx *gin.Context) {
	type Req struct {
		Address  string `json:"address,omitempty"`
		Birthday string `json:"birthday,omitempty"`
	}
	var req Req
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zap.S().Errorf("[addAddress] invalid params: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idStr, _ := ctx.Get("userID")
	userID, ok := idStr.(int64)
	if !ok {
		zap.S().Errorf("[addAddress] get userID failed")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "系统错误"})
		return
	}
	_, err := u.userClient.UpdateUser(ctx, &userPb.UpdateUserRequest{
		Id:       userID,
		Address:  req.Address,
		Birthday: req.Birthday,
	})
	if err != nil {
		zap.S().Errorf("[addAddress] update user failed: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "更新信息成功",
	})
}
