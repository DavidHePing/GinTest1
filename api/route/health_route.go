package route

import "github.com/gin-gonic/gin"

func HealthRouter(router *gin.RouterGroup) {
	router.GET("/health", func(ctx *gin.Context) {
		ctx.String(200, "healthy")
	})
}
