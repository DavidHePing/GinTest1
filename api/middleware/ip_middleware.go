package middleware

import "github.com/gin-gonic/gin"

func IpMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Request.Header.Add("X-Request-Id", "123.123.123")
		ctx.Next()
	}
}
