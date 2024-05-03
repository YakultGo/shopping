package jwt

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type LoginJWTMiddlewareBuilder struct {
	paths map[string]bool
	Handler
}

func NewLoginJWTMiddlewareBuilder(jwtHandler Handler) *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{
		paths:   make(map[string]bool),
		Handler: jwtHandler,
	}
}

func (l *LoginJWTMiddlewareBuilder) IgnorePath(path string) *LoginJWTMiddlewareBuilder {
	l.paths[path] = true
	return l
}

func (l *LoginJWTMiddlewareBuilder) Build() gin.HandlerFunc {
	// 注册time.Time类型，否则session无法存储
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		// 不需要登录校验
		if ok, val := l.paths[ctx.Request.URL.Path]; ok && val == true {
			return
		}
		// 使用JWT校验
		tokenStr := l.Handler.ExtractToken(ctx)
		claims := &UserClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return AccessTokenKey, nil
		})
		if err != nil {
			ctx.String(http.StatusUnauthorized, "未登录")
			zap.S().Errorf("解析token失败: %v", err)
			ctx.Abort()
			return
		}
		if !token.Valid {
			ctx.String(http.StatusUnauthorized, "未登录")
			zap.S().Error("token无效")
			ctx.Abort()
			return
		}
		if claims.UserAgent != ctx.Request.UserAgent() {
			ctx.String(http.StatusUnauthorized, "恶意入侵")
			ctx.Abort()
			return
		}
		err = l.CheckSession(ctx, claims.Ssid)
		if err != nil {
			zap.S().Errorf("session校验失败: %v", err)
			ctx.String(http.StatusUnauthorized, "未登录")
			ctx.Abort()
			return
		}
		ctx.Set("claims", claims)
		ctx.Set("userID", claims.UserId)
	}
}
