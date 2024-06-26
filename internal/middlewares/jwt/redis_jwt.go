package jwt

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

var (
	AccessTokenKey = []byte("4a1LwMzFjaCW4HrJETQsR8ybdYq82WMV")
)

type RedisJWTHandler struct {
	cmd redis.Cmdable
}

func NewRedisJWTHandler(cmd redis.Cmdable) Handler {
	return &RedisJWTHandler{
		cmd: cmd,
	}
}
func (h RedisJWTHandler) ExtractToken(ctx *gin.Context) string {
	// 使用JWT校验
	tokenHeader := ctx.GetHeader("Authorization")
	// Bearer token
	segments := strings.Split(tokenHeader, " ")
	if len(segments) != 2 {
		return ""
	}
	return segments[1]
}

func (h RedisJWTHandler) SetLoginToken(ctx *gin.Context, uid int64) error {
	ssid := uuid.New().String()
	err := h.SetJWTToken(ctx, uid, ssid)
	if err != nil {
		return err
	}
	return nil
}

func (h RedisJWTHandler) ClearToken(ctx *gin.Context) error {
	ctx.Header("x-jwt-token", "")
	c, _ := ctx.Get("claims")
	claims, ok := c.(*UserClaims)
	if !ok {
		return fmt.Errorf("系统错误")
	}
	err := h.cmd.Set(ctx, fmt.Sprintf("users:ssid:%s", claims.Ssid), 1,
		time.Hour*24*7).Err()
	if err != nil {
		return err
	}
	return nil
}

func (h RedisJWTHandler) CheckSession(ctx *gin.Context, ssid string) error {
	val, err := h.cmd.Exists(ctx, fmt.Sprintf("users:ssid:%s", ssid)).Result()
	switch {
	case errors.Is(err, redis.Nil):
		return nil
	case err == nil:
		if val == 0 {
			return nil
		}
		return errors.New("token 失效")
	default:
		return err
	}
}

func (h RedisJWTHandler) SetJWTToken(ctx *gin.Context, uid int64, ssid string) error {
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		},
		UserId:    uid,
		UserAgent: ctx.Request.UserAgent(),
		Ssid:      ssid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString(AccessTokenKey)
	if err != nil {
		return err
	}
	ctx.Header("x-jwt-token", tokenStr)
	return nil
}

type UserClaims struct {
	jwt.RegisteredClaims
	UserId    int64
	UserAgent string
	Ssid      string
}
