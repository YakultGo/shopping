package middlewares

import (
	"github.com/gin-gonic/gin"
	"shopping/internal/middlewares/cors"
	"shopping/internal/middlewares/jwt"
)

func NewMiddlewares(hdl jwt.Handler) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		cors.CorsHandler(),
		jwt.NewLoginJWTMiddlewareBuilder(hdl).
			IgnorePath("/user/signUp").
			IgnorePath("/user/login").
			Build(),
	}
}
